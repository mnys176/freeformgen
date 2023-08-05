package directives

import (
	_ "embed"
	"go/ast"
	"go/parser"
	"go/token"
	"html/template"
	"os"
	"sort"
	"strings"

	"github.com/mnys176/freeformgen/globals"
	"golang.org/x/mod/modfile"
)

const directiveFunctionSuffix string = "Directive"

var (
	packageSourcePath string
	directory         bool
	help              bool
)

//go:embed template.tmpl
var rawTemplate string

func Execute() error {

	if len(packageSourcePath) == 0 {
		return globals.NoSourceError()
	}

	fset := token.NewFileSet()
	templateData := make([]Directive, 0)
	if !directory {
		// Process single file.
		astFile, err := parser.ParseFile(fset, packageSourcePath, nil, 0)
		if err != nil {
			return err
		}
		foundDirectives, err := processASTFile(packageSourcePath, astFile)
		if err != nil {
			return err
		}
		templateData = append(templateData, foundDirectives...)
	} else {
		// Process all files in the directory.
		astPackages, err := parser.ParseDir(fset, packageSourcePath, nil, 0)
		if err != nil {
			return err
		}
		for _, astPackage := range astPackages {
			for path, astFile := range astPackage.Files {
				foundDirectives, err := processASTFile(path, astFile)
				if err != nil {
					return err
				}
				templateData = append(templateData, foundDirectives...)
			}
		}
	}

	sort.Slice(templateData, func(i, j int) bool {
		return templateData[i].Name < templateData[j].Name
	})

	// Overwrite output file in the current directory.
	f, err := os.Create("./directives_gen.go")
	if err != nil {
		return err
	}
	defer f.Close()

	// Parse `go.mod` file to derive full imports.
	modSrc, err := os.ReadFile("./go.mod")
	if err != nil {
		return err
	}
	modFile, err := modfile.Parse("go.mod", modSrc, nil)
	if err != nil {
		return err
	}

	tmpl := template.Must(template.New(globals.DirectivesCommand).Parse(rawTemplate))
	return tmpl.Execute(f, DirectiveTemplate{
		Prompt:     strings.Join(os.Args, " "),
		ModPath:    modFile.Module.Mod.Path,
		Directives: templateData,
	})
}

func processASTFile(goSrcPath string, astFile *ast.File) ([]Directive, error) {
	// Filter out exported directive functions.
	ast.FilterFile(astFile, func(s string) bool {
		return strings.HasSuffix(s, directiveFunctionSuffix)
	})

	// Load source bytes for indexing the actual tokens.
	goSrc, err := os.ReadFile(goSrcPath)
	if err != nil {
		return nil, err
	}

	// Parse source functions using AST.
	directives := make([]Directive, 0)
	ast.Inspect(astFile, func(n ast.Node) bool {
		if decl, ok := n.(*ast.FuncDecl); ok {
			declFunc := decl.Name.Name
			declName, _ := strings.CutSuffix(declFunc, directiveFunctionSuffix)
			declName = strings.ToLower(string(declName[:1])) + declName[1:]

			// There is a chance these could be nil.
			var declParams, declResults []*ast.Field
			if params := decl.Type.Params; params != nil {
				declParams = decl.Type.Params.List
			}
			if results := decl.Type.Results; results != nil {
				declResults = decl.Type.Results.List
			}

			// Directives must return a payload and an optional error.
			if len(declResults) == 0 || len(declResults) > 2 {
				return false
			}
			if len(declResults) == 2 {
				if t, ok := declResults[1].Type.(*ast.Ident); !ok || t.Name != "error" {
					return false
				}
			}

			directive := Directive{
				Name:   declName,
				Func:   declFunc,
				Pkg:    astFile.Name.Name,
				Params: make([]DirectiveParam, 0, len(declParams)),
			}
			for _, param := range declParams {
				var paramName strings.Builder
				paramTypeBytes := goSrc[param.Type.Pos()-astFile.Pos() : param.Type.End()-astFile.Pos()]
				for _, pn := range param.Names {
					paramName.WriteString(pn.Name + ", ")
				}
				directive.Params = append(
					directive.Params,
					DirectiveParam{
						Name:     strings.TrimSuffix(paramName.String(), ", "),
						Type:     string(paramTypeBytes),
						Variadic: strings.HasPrefix(string(paramTypeBytes), "..."),
					},
				)
			}
			for _, result := range declResults {
				resultTypeBytes := goSrc[result.Type.Pos()-astFile.Pos() : result.Type.End()-astFile.Pos()]
				directive.Returns = append(directive.Returns, string(resultTypeBytes))
			}
			directives = append(directives, directive)
		}
		return true
	})
	return directives, nil
}

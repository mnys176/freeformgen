package main

import (
	_ "embed"
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"sort"
	"strings"
	"text/template"

	"golang.org/x/mod/modfile"
)

type DirectiveParam struct {
	Name     string
	Type     string
	Variadic bool
}

type DirectiveReturn struct {
	Name string
	Type string
}

type Directive struct {
	Name    string
	Func    string
	Pkg     string
	Params  []DirectiveParam
	Returns []string
}

func (d *Directive) OtherParams() []DirectiveParam {
	if len(d.Params) > 0 {
		return d.Params[:len(d.Params)-1]
	}
	return d.Params
}

func (d *Directive) LastParam() *DirectiveParam {
	if len(d.Params) > 0 {
		return &d.Params[len(d.Params)-1]
	}
	return nil
}

type DirectiveTemplate struct {
	Prompt     string
	ModPath    string
	Directives []Directive
}

func (dt DirectiveTemplate) Imports() []string {
	found := make(map[string]bool)
	for _, d := range dt.Directives {
		found[dt.ModPath+"/"+d.Pkg] = true
	}

	imports := make([]string, len(found))
	i := 0
	for f := range found {
		imports[i] = f
		i++
	}
	sort.Strings(imports)
	return imports
}

type directiveCommand struct {
	Path      string
	Directory bool
	Help      bool
}

func (dc directiveCommand) String() string {
	const template string = `Path     : %s
Directory: %t
Help     : %t`

	return fmt.Sprintf(
		template,
		dc.Path,
		dc.Directory,
		dc.Help,
	)
}

func (dc directiveCommand) Handle() error {
	if dc.Help {
		fmt.Println(directiveCommandUsage)
		return nil
	}
	if len(dc.Path) == 0 {
		return errors.New("no source provided")
	}

	fset := token.NewFileSet()
	directives := make([]Directive, 0)
	if !dc.Directory {
		// Process single file.
		astFile, err := parser.ParseFile(fset, dc.Path, nil, 0)
		if err != nil {
			return err
		}
		foundDirectives, err := processASTFile(dc.Path, astFile)
		if err != nil {
			return err
		}
		directives = append(directives, foundDirectives...)
	} else {
		// Process all files in the directory.
		astPackages, err := parser.ParseDir(fset, dc.Path, nil, 0)
		if err != nil {
			return err
		}
		for _, astPackage := range astPackages {
			for path, astFile := range astPackage.Files {
				foundDirectives, err := processASTFile(path, astFile)
				if err != nil {
					return err
				}
				directives = append(directives, foundDirectives...)
			}
		}
	}

	sort.Slice(directives, func(i, j int) bool {
		return directives[i].Name < directives[j].Name
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

	tmpl := template.Must(template.New("directives").Parse(rawDirectivesTmpl))
	return tmpl.Execute(f, DirectiveTemplate{
		Prompt:     strings.Join(os.Args, " "),
		ModPath:    modFile.Module.Mod.Path,
		Directives: directives,
	})
}

const directiveFunctionSuffix string = "Directive"

//go:embed usage/directives-usage.txt
var directiveCommandUsage string

//go:embed templates/directives.tmpl
var rawDirectivesTmpl string

func parseDirectiveCommand(input []string) (*directiveCommand, error) {
	// Handle help.
	if len(input) == 1 || len(input) == 2 && (input[1] == "-h" || input[1] == "--help") {
		return &directiveCommand{Help: true}, nil
	}

	// Only `--help` option is valid without arguments.
	if len(input) == 2 && input[1] != "-h" && input[1] != "--help" && strings.HasPrefix(input[1], "-") {
		return nil, fmt.Errorf("unknown option: `%s`", input[1])
	}

	parsedDirectiveCommand := directiveCommand{Path: input[len(input)-1]}

	// Check if default behavior is desired (no options).
	if len(input) == 2 {
		return &parsedDirectiveCommand, nil
	}

	// var addNext bool
	// var previous string
	found := map[string]bool{"directory": false, "help": false}
	for _, token := range input[1 : len(input)-1] {
		// Add values to key-value pair options.
		// if addNext {
		// 	switch previous {
		// 	case "config":
		// 		absPath, _ := filepath.Abs(token)
		// 		parsedCreate.Config = absPath
		// 	}
		// 	addNext = false
		// 	continue
		// }

		switch token {
		case "-h", "--help":
			if !found["help"] {
				found["help"] = true
				parsedDirectiveCommand.Help = true
			}
		case "-d", "--directory":
			if !found["directory"] {
				found["directory"] = true
				parsedDirectiveCommand.Directory = true
			}
		// case "-c", "--config":
		// 	if !found["config"] {
		// 		found["config"] = true
		// 		previous = "config"
		// 		addNext = true
		// 	}
		default:
			return nil, fmt.Errorf("unknown option: `%s`", token)
		}
	}
	return &parsedDirectiveCommand, nil
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
				paramTypeBytes := goSrc[param.Type.Pos()-astFile.Pos() : param.Type.End()-astFile.Pos()]
				directive.Params = append(
					directive.Params,
					DirectiveParam{
						Name:     param.Names[0].Name,
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

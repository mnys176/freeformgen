package directives

import (
	_ "embed"
	"flag"
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

type directivesError struct {
	Message string
}

func (e directivesError) Error() string {
	return fmt.Sprintf("freeformgen: directives: %s", e.Message)
}

type DirectivesCommand struct {
	Path      string
	Directory bool
}

func (dc *DirectivesCommand) Execute() error {
	if len(dc.Path) == 0 {
		return noSourceProvidedError()
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

	tmpl := template.Must(template.New("directives").Parse(Template))
	return tmpl.Execute(f, DirectiveTemplate{
		Prompt:     strings.Join(os.Args, " "),
		ModPath:    modFile.Module.Mod.Path,
		Directives: directives,
	})
}

var DirectivesFlagSet = flag.NewFlagSet("directives", flag.ContinueOnError)

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

const directiveFunctionSuffix string = "Directive"

//go:embed usage.txt
var Usage string

//go:embed template.tmpl
var Template string

func noSourceProvidedError() error {
	return directivesError{"no source provided"}
}

func incorrectNumberOfArgumentsError(got int) error {
	return directivesError{fmt.Sprintf("wrong number of arguments: %d", got)}
}

func init() {
	DirectivesFlagSet.Usage = func() {
		// TODO: Dynamically generate this.
		fmt.Fprintln(os.Stdout, Usage)
	}
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

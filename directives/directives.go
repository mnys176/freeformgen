package directives

import (
	_ "embed"
	"sort"
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

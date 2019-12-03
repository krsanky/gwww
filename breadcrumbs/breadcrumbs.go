package breadcrumbs

import (
	"fmt"
	"html/template"

	"oldcode.org/repo/go/gow/lg"
)

type BCList struct {
	Bcs []BCItem
}

type BCItem struct {
	Name   string
	Path   string
	Active bool
}

func New() *BCList {
	return &BCList{}
}

var funcMap template.FuncMap

func init() {
	funcMap = template.FuncMap{
		"active": bc_active,
	}
}

func AddFuncs(t *template.Template) {
	t.Funcs(funcMap)
}

func bc_active(a bool) template.HTML {
	lg.Log.Printf("bc_active:%v", a)
	if a {
		return template.HTML(" active")
	} else {
		return template.HTML("")
	}
}

func (bc *BCItem) LI() template.HTML {
	if bc.Active {
		return template.HTML(fmt.Sprintf("<li class='breadcrumb-item active'>%s</li>",
			bc.Name))
	} else {
		return template.HTML(fmt.Sprintf("<li class='breadcrumb-item'><a href='%s'>%s</a> &gt;</li>",
			bc.Path, bc.Name))
	}
}

func (bcs *BCList) Append(name, path string) *BCList {
	bcs.Bcs = append(bcs.Bcs, BCItem{Name: name, Path: path, Active: false})
	return bcs
}

func (bcs *BCList) AppendActive(name string) *BCList {
	bcs.Bcs = append(bcs.Bcs, BCItem{Name: name, Active: true})
	return bcs
}

func (bcs *BCList) SetLastActive() {
	bcs.Bcs[len(bcs.Bcs)-1].Active = true
}

// Copyright (c) 2020 Paul Wisehart paul@oldcode.org
//
// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.

package breadcrumbs

import (
	"fmt"
	"html/template"

	"github.com/krsanky/gwww/lg"
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

func AddFuncs(fm template.FuncMap) {
	fm["active"] = bc_active
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

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

package main

import (
	"fmt"
	"html/template"
	"os"
)

var Fmap template.FuncMap

func f1() string {
	return "f1..."
}

func f2(s string) string {
	return "f2...|s:" + s
}

func init() {
	Fmap = template.FuncMap{
		"f1": f1,
		"f2": f2,
	}
}

func TmplTest() {
	fmt.Printf("Tmpl test...\n")

	tfiles := []string{
		"tmpl/test/t1.html",
		"tmpl/test/t2.html"}
	t := template.New("t1.html")
	t.Funcs(Fmap)
	_, err := t.ParseFiles(tfiles...)
	if err != nil {
		panic(err)
	}

	fmt.Printf("t.name:%s\n", t.Name())

	err = t.Execute(os.Stdout, nil)
	if err != nil {
		panic(err)
	}

}

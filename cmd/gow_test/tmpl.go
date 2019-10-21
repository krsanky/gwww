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

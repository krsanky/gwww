package tmpl

import (
	"fmt"
	"html/template"

	"oldcode.org/repo/go/gow/lg"
	krs_lorem "github.com/krsanky/lorem"
)

func AddFuncs(fm template.FuncMap) {
	fm["lorem"] = lorem
	fm["lorems"] = lorems
}

func lorem(a bool) template.HTML {
	lg.Log.Printf("lorem active:%v", a)
	return template.HTML(fmt.Sprintf("<h1>x active %v</h1>", a))
}

func lorems() template.HTML {
	return template.HTML(krs_lorem.Sentence())
}

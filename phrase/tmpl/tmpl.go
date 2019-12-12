package tmpl

import (
	"html/template"
	"strings"

	krs_lorem "github.com/krsanky/lorem"
)

func AddFuncs(fm template.FuncMap) {
	fm["lorem_w"] = lorem_w
	fm["lorem_ws"] = lorem_ws
	fm["lorem_s"] = lorem_s
	fm["lorem_p"] = lorem_p
	fm["lorem_ps"] = lorem_ps
}

func lorem_w() template.HTML {
	w := krs_lorem.Words(1, false)
	return template.HTML(w[0])
}

func lorem_ws(n int, r bool) template.HTML {
	ws := krs_lorem.Words(int(n), !r)
	out := strings.Join(ws, " ")
	return template.HTML(out)
}

func lorem_s() template.HTML {
	return template.HTML(krs_lorem.Sentence())
}

func lorem_p() template.HTML {
	return template.HTML(krs_lorem.Paragraph())
}

func lorem_ps(count int, common bool) template.HTML {
	ps := krs_lorem.Paragraphs(count, !common)
	out := strings.Join(ps, "\n\n")
	return template.HTML(out)
}

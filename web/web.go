package web

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"text/template"

	svg "github.com/ajstarks/svgo"
	lg "oldcode.org/gow/lg"
)

var tmpls map[string]*template.Template

func GetTmpls() map[string]*template.Template {
	return tmpls
}

func init() {
	tmpls = make(map[string]*template.Template)
}

func RenderPage(w http.ResponseWriter, page string, data interface{}) {
	lg.Log.Printf("RenderPage()...")
	headers := w.Header()
	headers.Add("Content-Type", "text/html")
	_, ok := tmpls[page]
	if !ok {
		dir, _ := os.Getwd()
		lg.Log.Printf("dir:%s", dir)
		os.Chdir("tmpl/")
		dir2, _ := os.Getwd()
		lg.Log.Printf("dir2:%s", dir2)
		ss := strings.Join([]string{page, "html"}, ".")
		lg.Log.Printf("ss:%s", ss)
		tmpls[page] = template.Must(template.ParseFiles("base.html", ss))
		os.Chdir(dir)
	}
	tmpls[page].Execute(w, data)
}

func Index(w http.ResponseWriter, r *http.Request) {
	headers := w.Header()
	headers.Add("Content-Type", "text/html")
	_, ok := tmpls["index"]
	if !ok {
		tmpls["index"] = template.Must(template.ParseFiles("tmpl/base.html", "tmpl/index.html"))
	}
	tmpls["index"].Execute(w, nil)
	//fmt.Fprintf(w, "<script>console.log('script');</script>\n")
}

func H1(w http.ResponseWriter, r *http.Request) {
	headers := w.Header()
	headers.Add("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1>H1</h1>\n")
}

func Other(w http.ResponseWriter, r *http.Request) {
	_, ok := tmpls["other"]
	if !ok {
		tmpls["other"] = template.Must(
			template.ParseFiles("tmpl/base.html", "tmpl/other.html"))
	}
	headers := w.Header()
	headers.Add("Content-Type", "text/html")
	//tmpl["other.html"].ExecuteTemplate(w, "name-of-define-block", nil)
	tmpls["other"].Execute(w, nil)
}

func Page3(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title1 string
		Items  []string
	}{
		Title1: "My page",
		Items: []string{
			"My photos",
			"My blog",
		},
	}

	RenderPage(w, "page3", data)
}

func Circle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	s := svg.New(w)
	s.Start(500, 500)
	s.Circle(250, 250, 125, "fill:none;stroke:black")
	s.End()
}

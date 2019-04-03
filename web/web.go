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
	lg.Log.Printf("RenderPage(%s)...", page)
	headers := w.Header()
	headers.Add("Content-Type", "text/html")
	_, tmpl_exists := tmpls[page]
	if !tmpl_exists {
		dir, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		if e := os.Chdir("tmpl/"); e != nil {
			panic(e)
		}
		p := strings.Join([]string{page, "html"}, ".")
		targs := []string{
			"base.html", 
			"navbar.tmpl", 
			"leftnav.tmpl", 
			"js_includes.tmpl", 
			p}
		lg.Log.Printf("targs:%s", strings.Join(targs, ","))
		tmpls[page] = template.Must(template.ParseFiles(targs...))
		if  e := os.Chdir(dir); e != nil {
			panic(e)
		}
	}
	tmpls[page].Execute(w, data)
}

// these views don't need to be here, don't put more here

func Index(w http.ResponseWriter, r *http.Request) {
	RenderPage(w, "index", nil)
}

func H1(w http.ResponseWriter, r *http.Request) {
	headers := w.Header()
	headers.Add("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1>H1</h1>\n")
}

func Other(w http.ResponseWriter, r *http.Request) {
	RenderPage(w, "other", nil)
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

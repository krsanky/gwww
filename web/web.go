package web

import (
	"fmt"
	"net"
	"net/http"
	"net/http/fcgi"
	"strings"
	"text/template"

	svg "github.com/ajstarks/svgo"
)

var tmpls map[string]*template.Template

func GetTmpls() map[string]*template.Template {
	return tmpls
}

func init() {
	tmpls = make(map[string]*template.Template)
}

func RenderPage(page string, w http.ResponseWriter, data interface{}) {
	headers := w.Header()
	headers.Add("Content-Type", "text/html")
	_, ok := tmpls[page]
	if !ok {
		ss := []string{"tmpl/", page, ".html"}
		tmpls[page] = template.Must(template.ParseFiles("tmpl/base.html",
			strings.Join(ss, "")))
	}
	tmpls[page].Execute(w, data)
}

func index(w http.ResponseWriter, r *http.Request) {
	headers := w.Header()
	headers.Add("Content-Type", "text/html")
	_, ok := tmpls["index"]
	if !ok {
		tmpls["index"] = template.Must(template.ParseFiles("tmpl/base.html", "tmpl/index.html"))
	}
	tmpls["index"].Execute(w, nil)
	//fmt.Fprintf(w, "<script>console.log('script');</script>\n")
}

func h1(w http.ResponseWriter, r *http.Request) {
	headers := w.Header()
	headers.Add("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1>H1</h1>\n")
}

func other(w http.ResponseWriter, r *http.Request) {
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

func page3(w http.ResponseWriter, r *http.Request) {
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

	RenderPage("page3", w, data)
}

func circle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	s := svg.New(w)
	s.Start(500, 500)
	s.Circle(250, 250, 125, "fill:none;stroke:black")
	s.End()
}

func Serve() {
	listener, err := net.Listen("tcp", "127.0.0.1:8088")
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/h1", h1)
	mux.HandleFunc("/other", other)
	mux.HandleFunc("/page3", page3)
	mux.HandleFunc("/page4",
		func(w http.ResponseWriter, r *http.Request) { RenderPage("page4", w, nil) })
	mux.HandleFunc("/circle", circle)

	fcgi.Serve(listener, mux)
}

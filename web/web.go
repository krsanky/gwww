package web

import (
	"fmt"
	"net"
	"net/http"
	"net/http/fcgi"
	"strings"
	"text/template"
)

var tmpls map[string]*template.Template

func GetTmpls() map[string]*template.Template {
	return tmpls
}

func init() {
	tmpls = make(map[string]*template.Template)
}

func RenderPage(page string, w http.ResponseWriter) {
	headers := w.Header()
	headers.Add("Content-Type", "text/html")
	_, ok := tmpls[page]
	if !ok {
		ss := []string{"tmpl/", page, ".html"}
		tmpls[page] = template.Must(template.ParseFiles("tmpl/base.html",
			strings.Join(ss, "")))
	}
	tmpls[page].Execute(w, nil)
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
	RenderPage("page3", w)
}

func Serve() {
	listener, err := net.Listen("tcp", "127.0.0.1:8088")
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/", index)
	http.HandleFunc("/h1", h1)
	http.HandleFunc("/other", other)
	http.HandleFunc("/page3", page3)
	fcgi.Serve(listener, nil)
}

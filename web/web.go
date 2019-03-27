package web

import (
	"fmt"
	"net"
	"net/http"
	"net/http/fcgi"
	"text/template"
)

var tmpls map[string]*template.Template

func init() {
	tmpls = make(map[string]*template.Template)
}

func index(w http.ResponseWriter, r *http.Request) {
	headers := w.Header()
	headers.Add("Content-Type", "text/html")
	_, ok := tmpls["index"]
	if (!ok) {
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
	tmpl := make(map[string]*template.Template)
	tmpl["other.html"] = template.Must(template.ParseFiles("tmpl/base.html", "tmpl/other.html"))
	headers := w.Header()
	headers.Add("Content-Type", "text/html")
	//tmpl["other.html"].ExecuteTemplate(w, "name-of-define-block", nil)
	tmpl["other.html"].Execute(w, nil)
}

func Serve() {
	listener, err := net.Listen("tcp", "127.0.0.1:8088")
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/", index)
	http.HandleFunc("/h1", h1)
	http.HandleFunc("/other", other)
	fcgi.Serve(listener, nil)
}

package web

import (
	"fmt"
	"net"
	"net/http"
	"net/http/fcgi"
	"text/template"
)

func handler(w http.ResponseWriter, r *http.Request) {
	headers := w.Header()
	headers.Add("Content-Type", "text/html")
	fmt.Fprintf(w, "HHHHHHHHHHHHHHHHHHHHHHH %s ----!", r.URL.Path[1:])
}

func h1(w http.ResponseWriter, r *http.Request) {
	headers := w.Header()
	headers.Add("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1>H1</h1>\n")
}

//tmpl := make(map[string]*template.Template)
//tmpl["index.html"] = template.Must(template.ParseFiles("index.html", "base.html"))
//tmpl["index.html"].ExecuteTemplate(w, "base.html", data)
func other(w http.ResponseWriter, r *http.Request) {
	tmpl := make(map[string]*template.Template)
	tmpl["other.html"] = template.Must(template.ParseFiles("tmpl/other.html", "tmpl/base.html"))
	headers := w.Header()
	headers.Add("Content-Type", "text/html")
	//fmt.Fprintf(w, "<h1>other</h1>\n")
	tmpl["other.html"].ExecuteTemplate(w, "base.html", nil)
}

func Serve() {
	listener, err := net.Listen("tcp", "127.0.0.1:8088")
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/", handler)
	http.HandleFunc("/h1", h1)
	http.HandleFunc("/other", other)
	fcgi.Serve(listener, nil)
}

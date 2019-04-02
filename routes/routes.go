package routes

import (
	"net"
	"net/http"
	"net/http/fcgi"
	"os"

	"oldcode.org/gow/formstuff"
	"oldcode.org/gow/lg"
	v1 "oldcode.org/gow/v1"
	"oldcode.org/gow/views"
	"oldcode.org/gow/web"
)

func Serve() {
	listener, err := net.Listen("tcp", "127.0.0.1:8088")
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", web.Index)

	mux.HandleFunc("/h1", web.H1)
	mux.HandleFunc("/other", web.Other)
	mux.HandleFunc("/page3", web.Page3)
	mux.HandleFunc("/page4",
		func(w http.ResponseWriter, r *http.Request) { web.RenderPage(w, "page4", nil) })
	mux.HandleFunc("/circle", web.Circle)

	mux.HandleFunc("/v1/index", v1.Index)

	mux.HandleFunc("/items", views.Items)

	mux.HandleFunc("/formstuff/index", formstuff.Index)

	dir, _ := os.Getwd()
	lg.Log.Printf("pre fcgi.Serve() dir:%s", dir)
	fcgi.Serve(listener, mux)
}

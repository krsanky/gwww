package routes

import (
	"net"
	"net/http"
	"net/http/fcgi"

	v1 "oldcode.org/gow/v1"
	web "oldcode.org/gow/web"
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

	fcgi.Serve(listener, mux)
}

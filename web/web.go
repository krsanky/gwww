package web

import (
	"fmt"
	"net"
	"net/http"
	"net/http/fcgi"
)

func handler(w http.ResponseWriter, r *http.Request) {
	headers := w.Header()
	headers.Add("Content-Type", "text/html")
	fmt.Fprintf(w, "HHHHHHHHHHHHHHHHHHHHHHH %s ----!", r.URL.Path[1:])
}

func Main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8088")
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/", handler)
	fcgi.Serve(listener, nil)
}

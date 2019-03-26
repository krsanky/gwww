package web

import (
	"fmt"
	"net"
	"net/http"
	"net/http/fcgi"
)

func homeView(w http.ResponseWriter, r *http.Request) {
	headers := w.Header()
	headers.Add("Content-Type", "text/html")
	//r.ParseForm()

	//io.WriteString(w, fmt.Sprintln("<p>Auth OK</p>"))
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func Main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8088")
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/", homeView)
	fcgi.Serve(listener, nil)
}

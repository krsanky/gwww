package routes

import (
	"net"
	"net/http"
	"net/http/fcgi"
	"os"

	"oldcode.org/gow/formstuff"
	"oldcode.org/gow/lg"
	"oldcode.org/gow/nostyle"
	v1 "oldcode.org/gow/v1"
	"oldcode.org/gow/views"
	"oldcode.org/gow/web"
)

func setupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", views.Index)
	mux.HandleFunc("/page3", views.Page3)
	mux.HandleFunc("/page4",
		func(w http.ResponseWriter, r *http.Request) { web.RenderPage(w, "page4", nil) })
	mux.HandleFunc("/circle", views.Circle)

	mux.HandleFunc("/v1/index", v1.Index)

	mux.HandleFunc("/items", views.Items)
	mux.HandleFunc("/artists", views.Artists)

	mux.HandleFunc("/formstuff/index", formstuff.Index)

	mux.HandleFunc("/nostyle", nostyle.Index)
	mux.HandleFunc("/nostyle/2", nostyle.V2)

	return mux
}

func Serve() {
	listener, err := net.Listen("tcp", "127.0.0.1:8088")
	if err != nil {
		panic(err)
	}

	mux := setupRoutes()

	dir, _ := os.Getwd()
	lg.Log.Printf("pre fcgi.Serve() dir:%s", dir)

	fcgi.Serve(listener, mux)
}

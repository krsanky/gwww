package routes

// make HasRoutes interface ?

//This is the main routes.
//A package can supply its own AddRoutes() function.
//Look at server package for adding the routes.

import (
	"net/http"

	v1 "oldcode.org/repo/go/gow/v1"
	"oldcode.org/repo/go/gow/views"
)

func AddRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/resume", views.Resume)
	mux.HandleFunc("/resume/", views.Resume)
	mux.HandleFunc("/phoon", views.Phoon)
	mux.HandleFunc("/msg", views.Msg)
	mux.HandleFunc("/circle", views.Circle)
	mux.HandleFunc("/dm", views.DirectMsg)

	mux.HandleFunc("/v1/index", v1.Index)

	//mux.HandleFunc("/page4",
	//		func(w http.ResponseWriter, r *http.Request) { web.RenderPage(w, "page4", nil) })

	mux.HandleFunc("/", views.Index)
}

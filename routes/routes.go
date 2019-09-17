package routes

//This is the main routes.
//A package can supply its own AddRoutes() function.
//Look at server package for adding the routes.
//Must keep the urt paths distinct.

import (
	"net/http"

	v1 "oldcode.org/home/wise/repo/go/oldcode.org/gow/v1"
	"oldcode.org/home/wise/repo/go/oldcode.org/gow/views"
)

func AddRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", views.Index)
	mux.HandleFunc("/msg", views.Msg)
	mux.HandleFunc("/circle", views.Circle)
	mux.HandleFunc("/dm", views.DirectMsg)

	mux.HandleFunc("/v1/index", v1.Index)

	//mux.HandleFunc("/page4",
	//		func(w http.ResponseWriter, r *http.Request) { web.RenderPage(w, "page4", nil) })
}

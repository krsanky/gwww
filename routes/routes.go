package routes

//This is the main routes.
//A package can supply its own AddRoutes() function.
//Look at server package for adding the routes.

import (
	"net/http"

	"oldcode.org/home/wise/repo/go/oldcode.org/gow/ttown"
	v1 "oldcode.org/home/wise/repo/go/oldcode.org/gow/v1"
	"oldcode.org/home/wise/repo/go/oldcode.org/gow/views"
)

func AddRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/old", views.Index)
	mux.HandleFunc("/old-index", views.Index)
	mux.HandleFunc("/msg", views.Msg)
	mux.HandleFunc("/circle", views.Circle)
	mux.HandleFunc("/dm", views.DirectMsg)

	mux.HandleFunc("/v1/index", v1.Index)

	//mux.HandleFunc("/page4",
	//		func(w http.ResponseWriter, r *http.Request) { web.RenderPage(w, "page4", nil) })

	mux.HandleFunc("/", ttown.Index)
}

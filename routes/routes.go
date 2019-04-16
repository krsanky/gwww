package routes

import (
	"net/http"

	"oldcode.org/gow/formstuff"
	"oldcode.org/gow/nostyle"
	v1 "oldcode.org/gow/v1"
	"oldcode.org/gow/views"
	"oldcode.org/gow/web"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", views.Index)
	mux.HandleFunc("/page3", views.Page3)
	mux.HandleFunc("/page4",
		func(w http.ResponseWriter, r *http.Request) { web.RenderPage(w, "page4", nil) })
	mux.HandleFunc("/circle", views.Circle)

	mux.HandleFunc("/v1/index", v1.Index)

	mux.HandleFunc("/items", views.Items)
	mux.HandleFunc("/artists", views.Artists)
	mux.HandleFunc("/artist", views.Artist)
	mux.HandleFunc("/album", views.Album)
	mux.HandleFunc("/track", views.Track)

	mux.HandleFunc("/formstuff/index", formstuff.Index)

	mux.HandleFunc("/nostyle", nostyle.Index)
	mux.HandleFunc("/nostyle/2", nostyle.V2)

	return mux
}

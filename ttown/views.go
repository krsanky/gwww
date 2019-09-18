package ttown

import (
	"net/http"

	"oldcode.org/home/wise/repo/go/oldcode.org/gow/web"
)

func AddRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/ttown", Index)
}

func Index(w http.ResponseWriter, r *http.Request) {
	data, _ := web.TmplData(r)
	tmpls := []string{
		"ttown/ttbase.html",
		"ttown/index.html"}
	web.Render(w, data, tmpls...)
}


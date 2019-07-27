package geo

import (
	"net/http"

	"oldcode.org/gow/breadcrumbs"
	"oldcode.org/gow/web"
)

func AddRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/geo", Index)
	mux.HandleFunc("/geo/map", Map)
}

func Index(w http.ResponseWriter, r *http.Request) {
	data, _ := web.TmplData(r)
	data["breadcrumbs"] = breadcrumbs.New().Append("Home", "/").AppendActive("GEO")
	tmpls := []string{
		"base.html",
		"nav.tmpl",
		"breadcrumbs.tmpl",
		"geo/index.html"}
	web.Render(w, data, tmpls...)
}

func Map(w http.ResponseWriter, r *http.Request) {
	data, _ := web.TmplData(r)
	bcs := breadcrumbs.New().Append("Home", "/").Append("GEO", "/geo")
	bcs.AppendActive("Map")
	data["breadcrumbs"] = bcs
	tmpls := []string{
		"base.html",
		"nav.tmpl",
		"breadcrumbs.tmpl",
		"geo/map.html"}
	web.Render(w, data, tmpls...)
}

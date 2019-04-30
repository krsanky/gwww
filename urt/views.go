package urt

import (
	"net/http"

	"oldcode.org/gow/breadcrumbs"
	"oldcode.org/gow/web"
)

func AddRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/urt", Index)
	mux.HandleFunc("/urt/radio", Radio)
}

func Index(w http.ResponseWriter, r *http.Request) {
	data, _ := web.TmplData(r)
	data["breadcrumbs"] = breadcrumbs.New().Append("Home", "/").AppendActive("URT", "/urt")
	tmpls := []string{
		"base.html",
		"nav.tmpl",
		"breadcrumbs.tmpl",
		"urt/index.html"}
	web.Render(w, data, tmpls...)
}

func Radio(w http.ResponseWriter, r *http.Request) {
	data, _ := web.TmplData(r)
	bcs := breadcrumbs.New().Append("Home", "/")
	bcs.Append("URT", "/urt")
	bcs.AppendActive("Radio", "/urt/radio")
	data["breadcrumbs"] = bcs
	tmpls := []string{
		"base.html",
		"nav.tmpl",
		"breadcrumbs.tmpl",
		"urt/radio.html"}
	web.Render(w, data, tmpls...)
}

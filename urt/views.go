package urt

import (
	"net/http"

	"oldcode.org/gow/breadcrumbs"
	"oldcode.org/gow/views"
	"oldcode.org/gow/web"
)

func AddRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/urt", Index)
	mux.HandleFunc("/urt/radio", Radio)
	mux.HandleFunc("/urt/radio/key", RadioKey)
	mux.HandleFunc("/urt/servers", Servers)
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

	if "POST" == r.Method {
		views.ShowFormData(r)
	}

	tmpls := []string{
		"base.html",
		"nav.tmpl",
		"breadcrumbs.tmpl",
		"urt/numpad.tmpl",
		"urt/radio.html"}
	web.Render(w, data, tmpls...)
}

func RadioKey(w http.ResponseWriter, r *http.Request) {
	data, _ := web.TmplData(r)
	bcs := breadcrumbs.New().Append("Home", "/")
	bcs.Append("URT", "/urt")
	bcs.AppendActive("Radio Key", "/urt/radio/key")
	data["breadcrumbs"] = bcs
	tmpls := []string{
		"base.html",
		"nav.tmpl",
		"breadcrumbs.tmpl",
		"urt/radio_key.html"}
	web.Render(w, data, tmpls...)
}

func Servers(w http.ResponseWriter, r *http.Request) {
	data, _ := web.TmplData(r)
	bcs := breadcrumbs.New().Append("Home", "/")
	bcs.Append("URT", "/urt")
	bcs.AppendActive("URT Servers", "/urt/servers")
	data["breadcrumbs"] = bcs
	tmpls := []string{
		"base.html",
		"nav.tmpl",
		"breadcrumbs.tmpl",
		"urt/servers.html"}
	web.Render(w, data, tmpls...)
}


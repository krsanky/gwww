package views

import (
	"net/http"
	"strings"

	"oldcode.org/gow/breadcrumbs"
	"oldcode.org/gow/lg"
	"oldcode.org/gow/web"
)

var A_Z = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I",
	"J", "K", "L", "M", "N", "O", "P", "Q", "R",
	"S", "T", "U", "V", "W", "X", "Y", "Z"}

func ShowFormData(r *http.Request) {
	r.ParseForm()
	for k, v := range r.Form {
		lg.Log.Printf("k: %s v: %s", k, strings.Join(v, ""))
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		lg.Log.Printf("views.Index(): NOT FOUND %s", r.URL.Path)
		http.NotFound(w, r)
		return
	}
	data, _ := web.TmplData(r)
	data["breadcrumbs"] = breadcrumbs.New().AppendActive("Home", "/")
	tmpls := []string{
		"base.html",
		"nav.tmpl",
		"breadcrumbs.tmpl",
		"index.html"}
	web.Render(w, data, tmpls...)
}

func Msg(w http.ResponseWriter, r *http.Request) {
	data, _ := web.TmplData(r)
	tmpls := []string{
		"base.html",
		"nav.tmpl",
		"msg.html"}
	web.Render(w, data, tmpls...)
}

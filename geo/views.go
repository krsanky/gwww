package geo

import (
	"net/http"

	"oldcode.org/gow/breadcrumbs"
	"oldcode.org/gow/web"
)

func AddRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/geo", Index)
}

func Index(w http.ResponseWriter, r *http.Request) {
	data, _ := web.TmplData(r)
	data["breadcrumbs"] = breadcrumbs.New().Append("Home", "/").AppendActive("GEO", "/geo")
	tmpls := []string{
		"base.html",
		"nav.tmpl",
		"breadcrumbs.tmpl",
		"geo/index.html"}
	web.Render(w, data, tmpls...)
}

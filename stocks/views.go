package stocks

import (
	"net/http"

	"oldcode.org/repo/go/gow/breadcrumbs"
	"oldcode.org/repo/go/gow/web"
)

func AddRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/stocks/index", Index)
}

func Index(w http.ResponseWriter, r *http.Request) {
	data, _ := web.TmplData(r)
	bcs := breadcrumbs.New().Append("Home", "/").Append("Projects", "/projects")
	bcs.AppendActive("Stocks")
	data["breadcrumbs"] = bcs
	tmpls := []string{
		"base.html",
		"breadcrumbs.tmpl",
		"stocks/index.html"}
	web.Render(w, data, tmpls...)
}

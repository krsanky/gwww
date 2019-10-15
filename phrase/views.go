package phrase

import (
	"net/http"

	"oldcode.org/home/wise/repo/go/oldcode.org/gow/breadcrumbs"
	"oldcode.org/home/wise/repo/go/oldcode.org/gow/web"
)

func AddRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/phrase", Index)
}

func Index(w http.ResponseWriter, r *http.Request) {
	data, _ := web.TmplData(r)
	data["breadcrumbs"] = breadcrumbs.New().Append("Home", "/").AppendActive("Phrase")
	tmpls := []string{
		"base.html",
		"breadcrumbs.tmpl",
		"phrase/index.html"}
	web.Render(w, data, tmpls...)
}

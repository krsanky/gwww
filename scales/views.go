package scales

import (
	"net/http"

	"oldcode.org/home/wise/repo/go/oldcode.org/gow/web"
)

func AddRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/scales", Index)
}

func Index(w http.ResponseWriter, r *http.Request) {
	data, _ := web.TmplData(r)
	tmpls := []string{
		"ttown/base.html",
		"scales/index.html"}
	web.Render(w, data, tmpls...)
}

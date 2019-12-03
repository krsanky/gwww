package scales

import (
	"net/http"

	"oldcode.org/repo/go/gow/web"
)

func AddRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/scales", Index)
}

func Index(w http.ResponseWriter, r *http.Request) {
	data, _ := web.TmplData(r)
	tmpls := []string{
		"base.html",
		"scales/index.html"}
	web.Render(w, data, tmpls...)
}

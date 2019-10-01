package canv_thing

import (
	"net/http"

	"github.com/justinas/nosurf"
	"oldcode.org/home/wise/repo/go/oldcode.org/gow/web"
)

func AddRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/canv-thing", Index)
}

func Index(w http.ResponseWriter, r *http.Request) {
	data, _ := web.TmplData(r)
	data["token"] = nosurf.Token(r)
	tmpls := []string{
		"ttown/base.html",
		"canv_thing/index.html"}
	web.Render(w, data, tmpls...)
}

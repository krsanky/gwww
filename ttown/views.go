package ttown

import (
	"net/http"

	"oldcode.org/repo/go/gow/web"
)

func AddRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/ttown/msg", Msg)
}

func Msg(w http.ResponseWriter, r *http.Request) {
	data, _ := web.TmplData(r)
	q := r.URL.Query()
	msg := q.Get("m")
	if msg != "" {
		data["msg"] = msg
	}
	tmpls := []string{
		"base.html",
		"ttown/msg.html"}
	web.Render(w, data, tmpls...)
}

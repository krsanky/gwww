package ttown

import (
	"net/http"

	"oldcode.org/home/wise/repo/go/oldcode.org/gow/urt"
	"oldcode.org/home/wise/repo/go/oldcode.org/gow/web"
)

func AddRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/ttown/msg", Msg)
	mux.HandleFunc("/ttown/urtctf", Urtctf)
	//mux.HandleFunc("/ttown/music", music.Music)
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

func Urtctf(w http.ResponseWriter, r *http.Request) {
	data, _ := web.TmplData(r)
	urt.UrtCtf(data)
	tmpls := []string{
		"base.html",
		"ttown/urtctf.html"}
	web.Render(w, data, tmpls...)
}

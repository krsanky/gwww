package ttown

import (
	"net/http"

	"oldcode.org/home/wise/repo/go/oldcode.org/gow/lg"
	"oldcode.org/home/wise/repo/go/oldcode.org/gow/music"
	"oldcode.org/home/wise/repo/go/oldcode.org/gow/urt"
	"oldcode.org/home/wise/repo/go/oldcode.org/gow/web"
)

func AddRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/ttown", Index)
	mux.HandleFunc("/ttown/msg", Msg)
	mux.HandleFunc("/ttown/urtctf", Urtctf)
	mux.HandleFunc("/ttown/music", music.Music)
}

func Index(w http.ResponseWriter, r *http.Request) {
	lg.Log.Printf("ttown.Index() method:%s", r.Method)
	data, _ := web.TmplData(r)
	tmpls := []string{
		"ttown/base.html",
		"ttown/index.html"}
	web.Render(w, data, tmpls...)
}

func Msg(w http.ResponseWriter, r *http.Request) {
	data, _ := web.TmplData(r)
	q := r.URL.Query()
	msg := q.Get("m")
	if msg != "" {
		data["msg"] = msg
	}
	tmpls := []string{
		"ttown/base.html",
		"ttown/msg.html"}
	web.Render(w, data, tmpls...)
}

func Urtctf(w http.ResponseWriter, r *http.Request) {
	data, _ := web.TmplData(r)
	urt.UrtCtf(data)
	tmpls := []string{
		"ttown/base.html",
		"ttown/urtctf.html"}
	web.Render(w, data, tmpls...)
}

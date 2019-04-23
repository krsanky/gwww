package gofed

import (
	"net/http"

	"oldcode.org/gow/lg"
	"oldcode.org/gow/session"
	"oldcode.org/gow/web"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tmpls := []string{
		"gofed/base.html",
		"gofed/nav.tmpl",
		"gofed/index.html"}
	web.Render(w, nil, tmpls...)
}

func Music(w http.ResponseWriter, r *http.Request) {
	lg.Log.Printf("gofed.Music()....")
	session := session.Manager.Load(r)
	cnt, err := session.GetInt("dumby_sess_counter")
	lg.Log.Printf("cnt:%d", cnt)
	if err != nil {
		panic(err)
	}
	err = session.PutInt(w, "dumby_sess_counter", cnt+1)
	if err != nil {
		panic(err)
	}
	tmpls := []string{
		"gofed/base.html",
		"gofed/nav.tmpl",
		"gofed/music.html"}
	web.Render(w, nil, tmpls...)
}

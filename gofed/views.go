package gofed

import (
	"net/http"

	"oldcode.org/gow/lg"
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
	tmpls := []string{
		"gofed/base.html",
		"gofed/nav.tmpl",
		"gofed/music.html"}
	web.Render(w, nil, tmpls...)
}

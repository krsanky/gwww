package univ

import (
	"net/http"

	"oldcode.org/home/wise/repo/go/oldcode.org/gow/web"
)

func AddRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/univ", Index)
	mux.HandleFunc("/univ/p2", Page2)
}

func Index(w http.ResponseWriter, r *http.Request) {
	data, _ := web.TmplData(r)
	data["content"] = "asdasd as'd'asd"
	tmpls := []string{
		"univ/base.html",
		"univ/txt.html"}
	web.Render(w, data, tmpls...)
}

func Page2(w http.ResponseWriter, r *http.Request) {
	tmpls := []string{
		"univ/base.html",
		"univ/page2.html"}
	web.Render(w, nil, tmpls...)
}

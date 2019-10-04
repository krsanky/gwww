package univ

import (
	"net/http"

	"oldcode.org/home/wise/repo/go/oldcode.org/gow/web"
)

func AddRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/univ", Index)
	mux.HandleFunc("/univ/p2", Page2)
	mux.HandleFunc("/univ/p3", Page3)
}

func Index(w http.ResponseWriter, r *http.Request) {
	data, _ := web.TmplData(r)
	data["content"] = "This is content from the view code"
	tmpls := []string{
		"univ/base.html",
		"univ/index.html"}
	web.Render(w, data, tmpls...)
}

func Page2(w http.ResponseWriter, r *http.Request) {
	tmpls := []string{
		"univ/base.html",
		"univ/page2.html"}
	web.Render(w, nil, tmpls...)
}

func Page3(w http.ResponseWriter, r *http.Request) {
	tmpls := []string{
		"univ/base.html",
		"univ/page3.html"}
	web.Render(w, nil, tmpls...)
}

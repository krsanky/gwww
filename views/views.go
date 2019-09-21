package views

import (
	"net/http"
	"strings"

	"oldcode.org/home/wise/repo/go/oldcode.org/gow/breadcrumbs"
	"oldcode.org/home/wise/repo/go/oldcode.org/gow/lg"
	"oldcode.org/home/wise/repo/go/oldcode.org/gow/web"
)

var A_Z = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I",
	"J", "K", "L", "M", "N", "O", "P", "Q", "R",
	"S", "T", "U", "V", "W", "X", "Y", "Z"}

func LogFormData(r *http.Request) {
	r.ParseForm()
	for k, v := range r.Form {
		lg.Log.Printf("k: %s v: %s", k, strings.Join(v, ""))
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
//	if r.URL.Path != "/" {
//		lg.Log.Printf("views.Index(): NOT FOUND %s", r.URL.Path)
//		http.NotFound(w, r)
//		return
//	}
	tmpls := []string{
		"base.html",
		"nav.tmpl",
		"index.html"}
	web.Render(w, nil, tmpls...)
}

func Msg(w http.ResponseWriter, r *http.Request) {
	// look for something in session like "flash_msg" ?
	// read the msg param

	data, _ := web.TmplData(r)
	data["breadcrumbs"] = breadcrumbs.New().Append("Home", "/").AppendActive("Message")
	q := r.URL.Query()
	msg := q.Get("m")
	if msg != "" {
		data["msg"] = msg
	}

	tmpls := []string{
		"base.html",
		"nav.tmpl",
		"breadcrumbs.tmpl",
		"msg.html"}
	web.Render(w, data, tmpls...)
}

func DirectMsg(w http.ResponseWriter, r *http.Request) {

}

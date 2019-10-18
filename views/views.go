package views

import (
	"net/http"
	"os/exec"
	"strings"

	"oldcode.org/home/wise/repo/go/gow/breadcrumbs"
	"oldcode.org/home/wise/repo/go/gow/lg"
	"oldcode.org/home/wise/repo/go/gow/web"
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
	data, _ := web.TmplData(r)
	out, err := exec.Command("phoon").Output()
	if err != nil {
		lg.Log.Printf("ERR:%s", err.Error())
		data["phoon"] = "err"
	} else {
		data["phoon"] = string(out)
	}
	tmpls := []string{
		"base.html",
		"index.html"}
	web.Render(w, data, tmpls...)
}

func Links(w http.ResponseWriter, r *http.Request) {
	lg.Log.Printf("ttown.Index() method:%s", r.Method)
	data, _ := web.TmplData(r)
	tmpls := []string{
		"base.html",
		"links.html"}
	web.Render(w, data, tmpls...)
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

func Resume(w http.ResponseWriter, r *http.Request) {
	tmpls := []string{
		"base.html",
		"resume.html"}
	web.Render(w, nil, tmpls...)
}

func Phoon(w http.ResponseWriter, r *http.Request) {
	data, _ := web.TmplData(r)
	out, err := exec.Command("phoon").Output()
	if err != nil {
		lg.Log.Printf("ERR:%s", err.Error())
		data["phoon"] = "err"
	} else {
		data["phoon"] = string(out)
	}
	tmpls := []string{
		"base.html",
		"phoon.html"}
	web.Render(w, data, tmpls...)
}

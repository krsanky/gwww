// Copyright (c) 2020 Paul Wisehart paul@oldcode.org
//
// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.

package views

import (
	"net/http"
	"os/exec"
	"strings"

	svg "github.com/ajstarks/svgo"
	"github.com/krsanky/gwww/breadcrumbs"
	"github.com/krsanky/gwww/lg"
	"github.com/krsanky/gwww/web"
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
	tmpls := []string{
		"base.html",
		"index.html"}
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

func Projects(w http.ResponseWriter, r *http.Request) {
	data, _ := web.TmplData(r)
	bcs := breadcrumbs.New()
	bcs.Append("Home", "/")
	bcs.Append("Projects", "")
	bcs.SetLastActive()
	data["breadcrumbs"] = bcs
	tmpls := []string{
		"base.html",
		"breadcrumbs.tmpl",
		"projects.html"}
	web.Render(w, data, tmpls...)
}

func Circle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	s := svg.New(w)
	s.Start(500, 500)
	s.Circle(250, 250, 125, "fill:none;stroke:black")
	s.End()
}

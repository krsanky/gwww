package phrase

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/schema"
	"github.com/justinas/nosurf"
	"oldcode.org/repo/go/gow/breadcrumbs"
	"oldcode.org/repo/go/gow/lg"
	"oldcode.org/repo/go/gow/web"
)

func AddRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/phrase/edit", Edit)
	mux.HandleFunc("/phrase/lorem", Lorem)
	mux.HandleFunc("/phrase/list", Phrases)
	mux.HandleFunc("/phrase/new", New)
	mux.HandleFunc("/phrase", Index)
}

func Index(w http.ResponseWriter, r *http.Request) {
	data, _ := web.TmplData(r)
	bcs := breadcrumbs.New().Append("Home", "/").AppendActive("Phrase")
	data["breadcrumbs"] = bcs
	tmpls := []string{
		"base.html",
		"breadcrumbs.tmpl",
		"phrase/index.html"}
	web.Render(w, data, tmpls...)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	data, _ := web.TmplData(r)
	bcs := breadcrumbs.New().Append("Home", "/").Append("Phrase", "/phrase")
	bcs.AppendActive("Edit")
	data["breadcrumbs"] = bcs
	data["token"] = nosurf.Token(r)

	data["phrase"] = Phrase{}
	pid_ := r.URL.Query().Get("p")
	pid, _ := strconv.Atoi(pid_)
	lg.Log.Printf("phrase.Edit() p:%d", pid)
	phr, err := GetPhrase(pid)
	if err != nil {
		data["error"] = err.Error()
	} else {
		data["phrase"] = phr
	}

	tmpls := []string{
		"base.html",
		"breadcrumbs.tmpl",
		"phrase/phrase.html"}
	web.Render(w, data, tmpls...)
}

func New(w http.ResponseWriter, r *http.Request) {
	data, _ := web.TmplData(r)
	bcs := breadcrumbs.New().Append("Home", "/").Append("Phrase", "/phrase")

	if strings.HasSuffix(r.URL.Path, "new") {
		bcs.AppendActive("New")
	} else {
		bcs.AppendActive("Edit")
	}

	data["breadcrumbs"] = bcs
	data["token"] = nosurf.Token(r)

	if "POST" == r.Method {
		err := r.ParseForm()
		if err != nil {
			lg.Log.Printf("ERR:%v", err)
			data["error"] = fmt.Sprintf("ERR:%v", err)
			goto Render
		}
		phrase := &Phrase{}
		decoder := schema.NewDecoder()
		decoder.IgnoreUnknownKeys(true)
		err = decoder.Decode(phrase, r.PostForm)
		if err != nil {
			lg.Log.Printf("ERR:%v", err)
			data["error"] = fmt.Sprintf("ERR schema:%v", err)
			goto Render
		}
		err = phrase.Insert()
		if err != nil {
			data["error"] = fmt.Sprintf("ERR insert:%v", err)
			goto Render
		}
		http.Redirect(w, r, "/msg?m=phrase+inserted", 303)
		return
	}

Render:
	tmpls := []string{
		"base.html",
		"breadcrumbs.tmpl",
		"phrase/phrase.html"}
	web.Render(w, data, tmpls...)
}

func Phrases(w http.ResponseWriter, r *http.Request) {
	data, _ := web.TmplData(r)
	bcs := breadcrumbs.New().Append("Home", "/").Append("Phrase", "/phrase")
	bcs.AppendActive("List")
	data["breadcrumbs"] = bcs
	data["token"] = nosurf.Token(r)

	if "POST" == r.Method {
		pathpre := r.FormValue("pathpre")
		lg.Log.Printf("pathpre[%s] len:%d", pathpre, len(pathpre))
	}

	ps, err := GetPhrases()
	if err != nil {
		data["error"] = fmt.Sprintf("ERR GetPhrases():%s", err.Error())
	} else {
		data["phrases"] = ps
	}

	tmpls := []string{
		"base.html",
		"breadcrumbs.tmpl",
		"phrase/phrases.html"}
	web.Render(w, data, tmpls...)
}

func Lorem(w http.ResponseWriter, r *http.Request) {
	data, _ := web.TmplData(r)

	bcs := breadcrumbs.New().Append("Home", "/").Append("Phrase", "/phrase")
	bcs.AppendActive("Lorem")
	data["breadcrumbs"] = bcs

	data["token"] = nosurf.Token(r)

	if "POST" == r.Method {
		pathpre := r.FormValue("pathpre")
		lg.Log.Printf("pathpre[%s] len:%d", pathpre, len(pathpre))
	}

	tmpls := []string{
		"base.html",
		"breadcrumbs.tmpl",
		"phrase/lorem.html"}
	web.Render(w, data, tmpls...)
}

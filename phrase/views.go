package phrase

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/schema"
	"github.com/justinas/nosurf"
	"oldcode.org/repo/go/gow/breadcrumbs"
	"oldcode.org/repo/go/gow/lg"
	"oldcode.org/repo/go/gow/web"
)

func AddRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/phrase", Index)
	mux.HandleFunc("/phrase/new", New)
	mux.HandleFunc("/phrase/edit", Edit)
	mux.HandleFunc("/phrase/edit/handle", Handler)
	mux.HandleFunc("/phrase/list", Phrases)
	mux.HandleFunc("/phrase/lorem", Lorem)
}

func Index(w http.ResponseWriter, r *http.Request) {
	data, _ := web.TmplData(r)
	bcs := breadcrumbs.New().Append("Home", "/").Append("Projects", "/projects")
	bcs.AppendActive("Phrase")
	data["breadcrumbs"] = bcs
	tmpls := []string{
		"base.html",
		"breadcrumbs.tmpl",
		"phrase/index.html"}
	web.Render(w, data, tmpls...)
}

func New(w http.ResponseWriter, r *http.Request) {
	data, _ := web.TmplData(r)
	bcs := breadcrumbs.New().Append("Home", "/")
	bcs.Append("Projects", "/projects")
	bcs.Append("Phrase", "/phrase")
	bcs.AppendActive("New")
	data["breadcrumbs"] = bcs
	data["token"] = nosurf.Token(r)
	tmpls := []string{
		"base.html",
		"breadcrumbs.tmpl",
		"phrase/phrase.html"}
	web.Render(w, data, tmpls...)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	data, _ := web.TmplData(r)
	bcs := breadcrumbs.New().Append("Home", "/")
	bcs.Append("Projects", "/projects")
	bcs.Append("Phrase", "/phrase")
	bcs.AppendActive("Edit")
	data["breadcrumbs"] = bcs
	data["token"] = nosurf.Token(r)

	pid_ := r.URL.Query().Get("p")
	pid, _ := strconv.Atoi(pid_)
	lg.Log.Printf("phrase.Edit() p:%d", pid)
	phrase, err := GetPhrase(pid)
	if err != nil {
		data["error"] = err.Error()
	} else {
		data["phrase"] = phrase
	}

	tmpls := []string{
		"base.html",
		"breadcrumbs.tmpl",
		"phrase/phrase.html"}
	web.Render(w, data, tmpls...)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if "POST" != r.Method {
		return
	}
	err := r.ParseForm()
	if err != nil {
		lg.Log.Printf("ERR:%v", err)
		return
	}
	phrase := new(Phrase)
	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)
	lg.Log.Printf("pf:%v", r.PostForm)
	err = decoder.Decode(phrase, r.PostForm)
	if err != nil {
		lg.Log.Printf("ERR:%v", err)
	}
	lg.Log.Printf("phrase.Id:%d path:%s", phrase.Id, phrase.Path)
	//if phrase.Id == 0 //NEW else EDIT

	//NEW
	phrase.Insert()
	http.Redirect(w, r, "/msg?m=new-phrase-inserted", 303)
}

func Phrases(w http.ResponseWriter, r *http.Request) {
	data, _ := web.TmplData(r)
	bcs := breadcrumbs.New().Append("Home", "/")
	bcs.Append("Projects", "/projects")
	bcs.Append("Phrase", "/phrase")
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

	bcs := breadcrumbs.New().Append("Home", "/")
	bcs.Append("Projects", "/projects")
	bcs.Append("Phrase", "/phrase")
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

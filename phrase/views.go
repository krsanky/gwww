package phrase

import (
	"fmt"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/justinas/nosurf"
	"oldcode.org/home/wise/repo/go/gow/breadcrumbs"
	"oldcode.org/home/wise/repo/go/gow/lg"
	"oldcode.org/home/wise/repo/go/gow/web"
)

func AddRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/phrase", PhraseView)
	mux.HandleFunc("/phrases", Phrases)
}

func PhraseView(w http.ResponseWriter, r *http.Request) {
	data, _ := web.TmplData(r)
	data["breadcrumbs"] = breadcrumbs.New().Append("Home", "/").AppendActive("Phrase")
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
	data["breadcrumbs"] = breadcrumbs.New().Append("Home", "/").AppendActive("Phrases")

	ps, err := GetPhrases()
	if err != nil {
		data["phrases"] = ps
	} else {
		//data["error"] = fmt.Sprintf("ERR GetPhrases():%s", err.Error())
	}

	tmpls := []string{
		"base.html",
		"breadcrumbs.tmpl",
		"phrase/phrases.html"}
	web.Render(w, data, tmpls...)
}

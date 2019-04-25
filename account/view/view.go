package view

import (
	"net/http"

	"github.com/justinas/nosurf"
	"oldcode.org/gow/account"
	"oldcode.org/gow/lg"
	"oldcode.org/gow/session"
	"oldcode.org/gow/web"
)

func Login(w http.ResponseWriter, r *http.Request) {
	lg.Log.Printf("LoginPage() method:%s", r.Method)
	data := make(map[string]interface{})

	if "POST" == r.Method {
		//view.ShowFormData(r)
		username := r.FormValue("inputEmail")
		password := r.FormValue("inputPassword")
		lg.Log.Printf("username:%s password:%s", username, password)

		ok := account.AuthUser(w, r, username, password)
		if ok {
			http.Redirect(w, r, "/", 303)
			return
		} else {
			data["error"] = "No match"
		}
	}

	data["token"] = nosurf.Token(r)
	lg.Log.Printf("LoginPage() token[%s]", data["token"])

	//view.Render(w, r, "account/login.html", ctx)
	tmpls := []string{
		"gofed/base.html",
		"gofed/nav.tmpl",
		"account/login.html"}
	web.Render(w, data, tmpls...)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session.Manager.Load(r).Clear(w)
	http.Redirect(w, r, "/", 303)
}

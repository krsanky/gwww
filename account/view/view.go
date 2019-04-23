package view

import (
	"net/http"

	"oldcode.org/gow/account"
	"oldcode.org/gow/lg"
	"oldcode.org/gow/session"
	"oldcode.org/gow/web"
)

func LoginPage(w http.ResponseWriter, r *http.Request) {
	lg.Log.Printf("LoginPage() method:%s", r.Method)
	data := make(map[string]interface{})

	if "POST" == r.Method {
		//view.ShowFormData(r)
		username := r.FormValue("username")
		password := r.FormValue("password")
		lg.Log.Printf("username:%s password:%s", username, password)

		ok := account.AuthUser(w, r, username, password)
		if ok {
			http.Redirect(w, r, "/", 303)
			return
		} else {
			data["error"] = "No match"
		}
	}

	//data["token"] = nosurf.Token(r)

	//view.Render(w, r, "account/login.html", ctx)
	tmpls := []string{
		"gofed/base.html",
		"gofed/nav.tmpl",
		"account/login.html"}
	web.Render(w, data, tmpls...)
}

func LogoutPage(w http.ResponseWriter, r *http.Request) {
	session.Manager.Load(r).Clear(w)
	http.Redirect(w, r, "/", 303)
}

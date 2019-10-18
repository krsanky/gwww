package view

import (
	"net/http"

	"github.com/justinas/nosurf"
	"oldcode.org/home/wise/repo/go/gow/account"
	"oldcode.org/home/wise/repo/go/gow/email"
	"oldcode.org/home/wise/repo/go/gow/lg"
	"oldcode.org/home/wise/repo/go/gow/views"
	"oldcode.org/home/wise/repo/go/gow/web"
)

func AddRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/account/login", Login)
	mux.HandleFunc("/account/login2", Login2)
	mux.HandleFunc("/account/logout", Logout)
	mux.HandleFunc("/account/password/reset", Reset)
}

func Reset(w http.ResponseWriter, r *http.Request) {
	data, _ := web.TmplData(r)
	tmpls := []string{
		"base.html",
		"nav.tmpl",
		"account/reset.html"}
	web.Render(w, data, tmpls...)
}

func Login(w http.ResponseWriter, r *http.Request) {
	lg.Log.Printf("LoginPage() method:%s", r.Method)
	data, _ := web.TmplData(r)

	if "POST" == r.Method {
		views.LogFormData(r)

		rbutton := r.FormValue("register_button")
		if rbutton == "register" {
			HandleRegister(w, r)
			return
		}

		email := r.FormValue("emailForLogin")
		password := r.FormValue("password")
		lg.Log.Printf("rbutton:%s email:%s password:%s", rbutton, email, password)

		ok := account.AuthUser(w, r, email, password)
		if ok {
			http.Redirect(w, r, "/msg?m=loggedin", 303)
			return
		} else {
			data["error"] = "No match"
		}
	}

	data["token"] = nosurf.Token(r)
	lg.Log.Printf("LoginPage() token[%s]", data["token"])

	tmpls := []string{
		"base.html",
		"nav.tmpl",
		"account/login.html"}
	web.Render(w, data, tmpls...)
}

func Login2(w http.ResponseWriter, r *http.Request) {
	lg.Log.Printf("account.Login2() method:%s", r.Method)
	data, _ := web.TmplData(r)

	if "POST" == r.Method {
		lg.Log.Printf("POST")
		views.LogFormData(r)

		username := r.FormValue("username")
		password := r.FormValue("password")
		lg.Log.Printf("username:%s password:%s", username, password)

		ok := account.AuthUser(w, r, username, password)
		if ok {
			http.Redirect(w, r, "/ttown/msg?m=loggedin", 303)
			return
		} else {
			data["error"] = "No match"
		}
	}

	data["token"] = nosurf.Token(r)
	lg.Log.Printf("LoginPage() token[%s]", data["token"])

	tmpls := []string{
		"base.html",
		"account/login2.html"}
	web.Render(w, data, tmpls...)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	account.Logout(w, r)
	http.Redirect(w, r, "/ttown/msg?m=loggedout", 303)
}

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	lg.Log.Printf("account.view.HandleRegister(): emailForReg:%s", r.FormValue("emailForReg"))
	_email := r.FormValue("emailForReg")

	//check that the email is unused
	//create a locked out account with that email

	var u *account.User
	var err error
	u, err = account.GetUserByEmail(_email)
	if err != nil {
		panic(err)
	}
	if u == nil {
		u = &account.User{}
		u.Email = _email
		err = u.SaveNew()
		if err != nil {
			lg.Log.Printf("ERR:%s", err)
		}
	} else {
		lg.Log.Printf("user found; email:%s", u.Email)
		http.Redirect(w, r, "/msg?m=emailinuse", 303)
		return
	}

	//create a

	// put this in a channel thing (fire and forget)
	go email.Send_test()

	http.Redirect(w, r, "/msg?m=checkemail", 303)
}

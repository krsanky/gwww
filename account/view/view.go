package view

import (
	"net/http"

	"github.com/justinas/nosurf"
	"oldcode.org/gow/account"
	"oldcode.org/gow/email"
	"oldcode.org/gow/lg"
	"oldcode.org/gow/views"
	"oldcode.org/gow/web"
)

func AddRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/account/login", Login)
	mux.HandleFunc("/account/logout", Logout)
}

func register(email string) {
	var u *account.User
	var err error
	u, err = account.GetUserByEmail(email)
	if err != nil {
	} else {
		u.Register()
	}
}
func HandleRegister(w http.ResponseWriter, r *http.Request) {
	//http.Redirect(w, r, "/msg?m=checkemail", 303)
	lg.Log.Printf("emailForReg:%s", r.FormValue("emailForReg"))
	email.Send_test()
}

func Login(w http.ResponseWriter, r *http.Request) {
	lg.Log.Printf("LoginPage() method:%s", r.Method)
	data, _ := web.TmplData(r)

	if "POST" == r.Method {
		views.ShowFormData(r)

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

func Logout(w http.ResponseWriter, r *http.Request) {
	account.Logout(w, r)
	http.Redirect(w, r, "/?msg=loggedout", 303)
}

func PostReg(w http.ResponseWriter, r *http.Request) {

}

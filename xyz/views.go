package xyz

import (
	"net/http"
	"strconv"

	"oldcode.org/gow/account"
	"oldcode.org/gow/lg"
	"oldcode.org/gow/web"
)

func Users(w http.ResponseWriter, r *http.Request) {
	lg.Log.Printf("xyz.Users() method:%s", r.Method)
	//data := make(map[string]interface{})

	data, err := web.TmplData(r)
	if err != nil {
		lg.Log.Printf("err1:%s", err)
	}

	users, err := account.GetUsers()
	if err != nil {
		lg.Log.Printf("err:%s", err)
	} else {
		data["users"] = users
		lg.Log.Printf("u1:%s", users[0])
	}

	tmpls := []string{
		"base.html",
		"nav.tmpl",
		"xyz/users.html"}
	web.Render(w, data, tmpls...)
}

func User(w http.ResponseWriter, r *http.Request) {
	lg.Log.Printf("xyz.User() method:%s", r.Method)
	data := make(map[string]interface{})
	q := r.URL.Query()

	uid, err := strconv.Atoi(q.Get("u"))
	if err != nil {
		lg.Log.Printf("xzy.User() err1:%s", err.Error())
	}
	user, err := account.GetUserById(uid)
	if err != nil {
		lg.Log.Printf("xzy.User() err2:%s", err.Error())
	}
	data["user"] = user

	tmpls := []string{
		"base.html",
		"nav.tmpl",
		"xyz/user.html"}
	web.Render(w, data, tmpls...)
}

func Become(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	uid, err := strconv.Atoi(q.Get("u"))
	if err != nil {
		lg.Log.Printf("xzy.Become() err1:%s", err.Error())
	}
	user, err := account.GetUserById(uid)
	if err != nil {
		lg.Log.Printf("xzy.Become() err2:%s", err.Error())
	}
	account.LoginUser(w, r, user)
	http.Redirect(w, r, "/msg?m=uloggedin", 303)
}

func AddRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/xyz/users", Users)
	mux.HandleFunc("/xyz/user", User)
	mux.HandleFunc("/xyz/become", Become)
}

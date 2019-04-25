package xyz

import (
	"net/http"

	"oldcode.org/gow/account"
	"oldcode.org/gow/lg"
	"oldcode.org/gow/web"
)

func Users(w http.ResponseWriter, r *http.Request) {
	lg.Log.Printf("xyz.Users() method:%s", r.Method)
	data := make(map[string]interface{})
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
	tmpls := []string{
		"base.html",
		"nav.tmpl",
		"xyz/user.html"}
	web.Render(w, data, tmpls...)
}

func AddRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/xyz/users", Users)
	mux.HandleFunc("/xyz/user/", User)
}

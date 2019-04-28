package xyz

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/schema"
	"github.com/justinas/nosurf"
	"oldcode.org/gow/account"
	"oldcode.org/gow/lg"
	"oldcode.org/gow/web"
)

func Users(w http.ResponseWriter, r *http.Request) {
	lg.Log.Printf("xyz.Users() method:%s", r.Method)

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
	data, _ := web.TmplData(r)
	q := r.URL.Query()

	uid, err := strconv.Atoi(q.Get("u"))
	if err != nil {
		lg.Log.Printf("xzy.User() err1:%s", err.Error())
	}
	user, err := account.GetUserById(uid)
	if err != nil {
		lg.Log.Printf("xzy.User() err2:%s", err.Error())
	}
	data["user1"] = user

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

func SendEmail(w http.ResponseWriter, r *http.Request) {
	tmpls := []string{
		"base.html",
		"nav.tmpl",
		"xyz/send_email.html"}
	web.Render(w, nil, tmpls...)
}

//type User struct {
//	Id         int
//	Password   string
//	Is_super   bool
//	Username   string
//	First_name string
//	Last_name  string
//	Email      string
//	Is_staff   bool
//	Is_active  bool
//r.PostForm (after r.ParseForm()
func AddUser(w http.ResponseWriter, r *http.Request) {
	user := &account.User{}
	data, _ := web.TmplData(r)
	data["token"] = nosurf.Token(r)
	data["user"] = user

	if "POST" == r.Method {
		decoder := schema.NewDecoder()
		decoder.IgnoreUnknownKeys(true)
		err := r.ParseForm()
		for k, v := range r.Form {
			lg.Log.Printf("k: %s v: %s", k, strings.Join(v, ""))
		}
		if err != nil {
			lg.Log.Printf("err:%s", err.Error())
			goto Render
		}
		err = decoder.Decode(user, r.PostForm)
		if err != nil {
			lg.Log.Printf("err:%s", err.Error())
			goto Render
		}
		lg.Log.Printf("user:%v", user)
		//http.Redirect(w, r, newUrl, http.StatusSeeOther)
	}

Render:
	tmpls := []string{
		"base.html",
		"nav.tmpl",
		"xyz/add_user.html"}
	web.Render(w, data, tmpls...)
}

func AddRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/xyz/users", Users)
	mux.HandleFunc("/xyz/user", User)
	mux.HandleFunc("/xyz/become", Become)
	mux.HandleFunc("/xyz/add-user", AddUser)
	mux.HandleFunc("/xyz/send-email", SendEmail)
}
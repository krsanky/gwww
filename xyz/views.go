package xyz

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/schema"
	"github.com/justinas/nosurf"
	"oldcode.org/home/wise/repo/go/oldcode.org/gow/account"
	"oldcode.org/home/wise/repo/go/oldcode.org/gow/breadcrumbs"
	"oldcode.org/home/wise/repo/go/oldcode.org/gow/lg"
	"oldcode.org/home/wise/repo/go/oldcode.org/gow/web"
)

func AddRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/xyz", Index)
	mux.HandleFunc("/xyz/users", Users)
	mux.HandleFunc("/xyz/user", User)
	mux.HandleFunc("/xyz/become", Become)
	mux.HandleFunc("/xyz/add-user", AddUser)
	mux.HandleFunc("/xyz/send-email", SendEmail)
	mux.HandleFunc("/xyz/colors", Colors)
}

func Index(w http.ResponseWriter, r *http.Request) {
	data, err := web.TmplData(r)
	if err != nil {
		lg.Log.Printf("err:%s", err)
	}
	bcs := breadcrumbs.New()
	bcs.Append("Home", "/")
	bcs.Append("XYZ", "/xyz")
	bcs.SetLastActive()
	data["breadcrumbs"] = bcs
	tmpls := []string{
		"base.html",
		"nav.tmpl",
		"breadcrumbs.tmpl",
		"xyz/index.html"}
	web.Render(w, data, tmpls...)
}

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
		//lg.Log.Printf("u1:%s", users[0])
	}

	bcs := breadcrumbs.New()
	bcs.Append("Home", "/")
	bcs.Append("XYZ", "/xyz")
	bcs.Append("Users", "/xyz/users")
	bcs.SetLastActive()
	data["breadcrumbs"] = bcs

	tmpls := []string{
		"base.html",
		"nav.tmpl",
		"breadcrumbs.tmpl",
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
	data["user"] = user

	bcs := breadcrumbs.New()
	bcs.Append("Home", "/")
	bcs.Append("XYZ", "/xyz")
	bcs.Append("User", "/xyz/user")
	bcs.SetLastActive()
	data["breadcrumbs"] = bcs
	tmpls := []string{
		"base.html",
		"nav.tmpl",
		"breadcrumbs.tmpl",
		"xyz/user_form.tmpl",
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
	bcs := breadcrumbs.New()
	bcs.Append("Home", "/")
	bcs.Append("XYZ", "/xyz")
	bcs.Append("Add User", "/xyz/add-user")
	bcs.SetLastActive()
	data["breadcrumbs"] = bcs
	tmpls := []string{
		"base.html",
		"nav.tmpl",
		"breadcrumbs.tmpl",
		"xyz/user_form.tmpl",
		"xyz/add_user.html"}
	web.Render(w, data, tmpls...)
}

func Colors(w http.ResponseWriter, r *http.Request) {
	data, _ := web.TmplData(r)
	bcs := breadcrumbs.New()
	bcs.Append("Home", "/")
	bcs.Append("XYZ", "/xyz")
	bcs.AppendActive("Colors")
	data["breadcrumbs"] = bcs
	tmpls := []string{
		"base.html",
		"nav.tmpl",
		"breadcrumbs.tmpl",
		"xyz/colors.html"}
	web.Render(w, data, tmpls...)
}

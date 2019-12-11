package xyz

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/schema"
	"github.com/justinas/nosurf"
	"oldcode.org/repo/go/gow/account"
	"oldcode.org/repo/go/gow/breadcrumbs"
	"oldcode.org/repo/go/gow/lg"
	"oldcode.org/repo/go/gow/secure"
	"oldcode.org/repo/go/gow/web"
)

func AddRoutes(mux *http.ServeMux) {
	//mux.Handle("/xyz", http.HandlerFunc(Index))
	mux.Handle("/xyz", secure.SuperOnly(Index))

	mux.Handle("/xyz/users", secure.SuperOnly(Users))
	mux.Handle("/xyz/user", secure.SuperOnly(User))
	mux.Handle("/xyz/become", secure.SuperOnly(Become))
	mux.Handle("/xyz/add-user", secure.SuperOnly(AddUser))
	mux.Handle("/xyz/send-email", secure.SuperOnly(SendEmail))
	mux.Handle("/xyz/colors", secure.SuperOnly(Colors))
	mux.Handle("/xyz/post-test", secure.SuperOnly(PostTest))
	mux.Handle("/xyz/tdquote", secure.SuperOnly(TDQuote))
	mux.Handle("/xyz/semantic-ui", secure.SuperOnly(SemanticUI))
}

func Index(w http.ResponseWriter, r *http.Request) {
	data, _ := web.TmplData(r)
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

func PostTest(w http.ResponseWriter, r *http.Request) {
	data, _ := web.TmplData(r)
	bcs := breadcrumbs.New()
	bcs.Append("Home", "/")
	bcs.Append("XYZ", "/xyz")
	bcs.AppendActive("Post Test")
	data["breadcrumbs"] = bcs
	tmpls := []string{
		"base.html",
		"nav.tmpl",
		"breadcrumbs.tmpl",
		"xyz/post_test.html"}
	web.Render(w, data, tmpls...)
}

func TDQuote(w http.ResponseWriter, r *http.Request) {
	lg.Log.Printf("xyz.TDQuote...")
	data, _ := web.TmplData(r)
	bcs := breadcrumbs.New()
	bcs.Append("Home", "/")
	bcs.Append("XYZ", "/xyz")
	bcs.Append("TD Quote", "/xyz/tdquote")
	bcs.SetLastActive()
	data["breadcrumbs"] = bcs
	tmpls := []string{
		"base.html",
		"nav.tmpl",
		"breadcrumbs.tmpl",
		"xyz/tdquote.html"}
	web.Render(w, data, tmpls...)
}

func SemanticUI(w http.ResponseWriter, r *http.Request) {
	data, _ := web.TmplData(r)
	bcs := breadcrumbs.New()
	bcs.Append("Home", "/")
	bcs.Append("XYZ", "/xyz")
	bcs.Append("Semantic UI Test", "")
	bcs.SetLastActive()
	data["breadcrumbs"] = bcs
	tmpls := []string{
		"base-sem.html",
		"nav.tmpl",
		"breadcrumbs.tmpl",
		"xyz/semantic-ui.html"}
	web.Render(w, data, tmpls...)
}

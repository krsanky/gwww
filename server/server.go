package server

import (
	"net"
	"net/http"
	"net/http/fcgi"
	"os"

	"github.com/justinas/nosurf"
	"oldcode.org/gow/account"
	account_view "oldcode.org/gow/account/view"
	"oldcode.org/gow/db"
	"oldcode.org/gow/geo"
	"oldcode.org/gow/lg"
	"oldcode.org/gow/music"
	"oldcode.org/gow/routes"
	"oldcode.org/gow/session"
	"oldcode.org/gow/urt"
	"oldcode.org/gow/xyz"
)

//try:
//https://github.com/alexedwards/stack
func Serve() {
	listener, err := net.Listen("tcp", "127.0.0.1:8088")
	if err != nil {
		panic(err)
	}

	dir, _ := os.Getwd()
	lg.Log.Printf("pre fcgi.Serve() dir:%s", dir)

	db.Open()
	db.InitDB()

	//mux is a handler, because ServeMux implements ServeHTTP()
	mux := http.NewServeMux()
	routes.AddRoutes(mux)
	account_view.AddRoutes(mux)
	xyz.AddRoutes(mux)
	music.AddRoutes(mux)
	urt.AddRoutes(mux)
	geo.AddRoutes(mux)

	// ORDER MATTERS ... acccount depends on session
	h := nosurf.NewPure(mux)
	//h = M1(h, "->h1")
	h = account.AddUser(h)
	session.Init()
	h = session.Session.LoadAndSave(h)
	fcgi.Serve(listener, h)
}

func M1(h http.HandlerFunc, extra_arg string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lg.Log.Printf("M1--%s--- %s", extra_arg, r.RequestURI)
		h.ServeHTTP(w, r)
	}
}

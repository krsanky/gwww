package server

import (
	"net"
	"net/http"
	"net/http/fcgi"

	"github.com/justinas/nosurf"
	"github.com/krsanky/gwww/account"
	account_view "github.com/krsanky/gwww/account/view"
	"github.com/krsanky/gwww/db"
	"github.com/krsanky/gwww/lg"
	"github.com/krsanky/gwww/routes"
	"github.com/krsanky/gwww/session"
	"github.com/krsanky/gwww/settings"
	"github.com/krsanky/gwww/xyz"
)

func setupRoutes() *http.ServeMux {
	//mux is a handler, because ServeMux implements ServeHTTP()
	mux := http.NewServeMux()

	routes.AddRoutes(mux)
	account_view.AddRoutes(mux)
	xyz.AddRoutes(mux)

	return mux
}

func Serve(sfile string) {
	settings.Init(sfile)

	db.Init()

	mux := setupRoutes()

	// ORDER MATTERS and it's kind of reversed
	h := nosurf.NewPure(mux)
	//h = M1(h, "->h1")
	//h = secure.HHHEnforceSuperUser(h)
	h = account.AddUser(h)
	session.Init()
	h = session.Session.LoadAndSave(h)

	listener, err := net.Listen("tcp", "127.0.0.1:8088")
	if err != nil {
		panic(err)
	}

	fcgi.Serve(listener, h)
}

func M1(h http.HandlerFunc, extra_arg string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lg.Log.Printf("M1--%s--- %s", extra_arg, r.RequestURI)
		h.ServeHTTP(w, r)
	}
}

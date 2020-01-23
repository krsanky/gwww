package server

import (
	"net"
	"net/http"
	"net/http/fcgi"

	"github.com/justinas/nosurf"
	"github.com/krsanky/gwww/account"
	"github.com/krsanky/gwww/db"
	"github.com/krsanky/gwww/lg"
	"github.com/krsanky/gwww/session"
	"github.com/krsanky/gwww/settings"
)

var mux *http.ServeMux

func init() {
	mux = http.NewServeMux()
}

func Handle(path string, h http.Handler) {
	mux.Handle(path, h)
}

func Serve(sfile string) {
	settings.Init(sfile)

	db.Init()

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

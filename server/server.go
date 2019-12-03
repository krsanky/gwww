package server

import (
	"net"
	"net/http"
	"net/http/fcgi"
	"os"

	"github.com/justinas/nosurf"
	"oldcode.org/home/wise/repo/go/oldcode.org/gow/account"
	account_view "oldcode.org/home/wise/repo/go/oldcode.org/gow/account/view"
	"oldcode.org/home/wise/repo/go/oldcode.org/gow/canv_thing"
	"oldcode.org/home/wise/repo/go/oldcode.org/gow/db"
	"oldcode.org/home/wise/repo/go/oldcode.org/gow/geo"
	"oldcode.org/home/wise/repo/go/oldcode.org/gow/lg"
	"oldcode.org/home/wise/repo/go/oldcode.org/gow/music"
	"oldcode.org/home/wise/repo/go/oldcode.org/gow/phrase"
	"oldcode.org/home/wise/repo/go/oldcode.org/gow/routes"
	"oldcode.org/home/wise/repo/go/oldcode.org/gow/scales"
	"oldcode.org/home/wise/repo/go/oldcode.org/gow/session"
	"oldcode.org/home/wise/repo/go/oldcode.org/gow/settings"
	"oldcode.org/home/wise/repo/go/oldcode.org/gow/ttown"
	"oldcode.org/home/wise/repo/go/oldcode.org/gow/univ"
	"oldcode.org/home/wise/repo/go/oldcode.org/gow/urt"
	"oldcode.org/home/wise/repo/go/oldcode.org/gow/xyz"
	"oldcode.org/home/wise/repo/go/oldcode.org/gow/zz"
)

//try:
//https://github.com/alexedwards/stack
func Serve(sfile string) {
	settings.Init(sfile)
	listener, err := net.Listen("tcp", "127.0.0.1:8088")
	if err != nil {
		panic(err)
	}

	// get rid of this dir stuff...
	dir, _ := os.Getwd()
	lg.Log.Printf("server.Serve() dir:%s", dir)
	os.Chdir("/home/wise/data/GO/gow")
	dir, _ = os.Getwd()
	lg.Log.Printf("--now dir:%s", dir)

	db.InitDB()

	//mux is a handler, because ServeMux implements ServeHTTP()
	mux := http.NewServeMux()
	routes.AddRoutes(mux)
	account_view.AddRoutes(mux)
	xyz.AddRoutes(mux)
	music.AddRoutes(mux)
	urt.AddRoutes(mux)
	geo.AddRoutes(mux)
	zz.AddRoutes(mux)
	ttown.AddRoutes(mux)
	canv_thing.AddRoutes(mux)
	univ.AddRoutes(mux)
	scales.AddRoutes(mux)
	phrase.AddRoutes(mux)

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

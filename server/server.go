package server

import (
	"net"
	"net/http"
	"net/http/fcgi"

	"github.com/justinas/nosurf"
	"oldcode.org/home/wise/repo/go/gow/account"
	account_view "oldcode.org/home/wise/repo/go/gow/account/view"
	"oldcode.org/home/wise/repo/go/gow/canv_thing"
	"oldcode.org/home/wise/repo/go/gow/db"
	"oldcode.org/home/wise/repo/go/gow/geo"
	"oldcode.org/home/wise/repo/go/gow/lg"
	"oldcode.org/home/wise/repo/go/gow/music"
	"oldcode.org/home/wise/repo/go/gow/phrase"
	"oldcode.org/home/wise/repo/go/gow/routes"
	"oldcode.org/home/wise/repo/go/gow/scales"
	"oldcode.org/home/wise/repo/go/gow/session"
	"oldcode.org/home/wise/repo/go/gow/settings"
	"oldcode.org/home/wise/repo/go/gow/ttown"
	"oldcode.org/home/wise/repo/go/gow/univ"
	"oldcode.org/home/wise/repo/go/gow/urt"
	"oldcode.org/home/wise/repo/go/gow/xyz"
	"oldcode.org/home/wise/repo/go/gow/zz"
)

//try:
//https://github.com/alexedwards/stack
func Serve(sfile string) {
	settings.Init(sfile)
	listener, err := net.Listen("tcp", "127.0.0.1:8088")
	if err != nil {
		panic(err)
	}

//	dir, _ := os.Getwd()
//	lg.Log.Printf("server.Serve() dir:%s", dir)
//	os.Chdir("/home/wise/data/GO/gow")
//	dir, _ = os.Getwd()
//	lg.Log.Printf("--now dir:%s", dir)
	db.InitDB() // This is kinda important

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

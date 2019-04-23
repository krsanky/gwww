package server

import (
	"net"
	"net/http"
	"net/http/fcgi"
	"os"

	"oldcode.org/gow/db"
	"oldcode.org/gow/lg"
	"oldcode.org/gow/routes"
	"oldcode.org/gow/session"
	account_view "oldcode.org/gow/account/view"
)

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

	session.Init()
	h := session.Manager.Use(mux)
	fcgi.Serve(listener, h)
}

//	// ORDER MATTERS ... acccount depends on session
//	h := nosurf.NewPure(routes.Router())
//	h = account.AddUser(h)
//	session.Init()
//	h = session.Manager.Use(h)

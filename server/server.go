package server

import (
	"net"
	"net/http/fcgi"
	"os"

	"oldcode.org/gow/db"
	"oldcode.org/gow/lg"
	"oldcode.org/gow/routes"
)

//	// ORDER MATTERS ... acccount depends on session
//	h := nosurf.NewPure(routes.Router())
//	h = account.AddUser(h)
//	session.Init()
//	h = session.Manager.Use(h)
func Serve() {
	listener, err := net.Listen("tcp", "127.0.0.1:8088")
	if err != nil {
		panic(err)
	}

	db.Open()

	mux := routes.SetupRoutes()

	dir, _ := os.Getwd()
	lg.Log.Printf("pre fcgi.Serve() dir:%s", dir)

	fcgi.Serve(listener, mux)
}

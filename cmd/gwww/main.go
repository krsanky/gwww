package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/krsanky/gwww/server"
	"github.com/krsanky/gwww/views"
)

func main() {
	fmt.Printf("num args:%d\n", len(os.Args))
	if len(os.Args) == 2 {
		settings := os.Args[1]
		server := server.NewServer()
		server.Handle("/", http.HandlerFunc(views.Index))
		server.Serve(settings)
	} else {
		usage()
	}
}

func usage() {
	fmt.Printf("\ngwww\n")
}

package main

import (
	"fmt"
	"go/build"
	"os"
	"strings"

	"oldcode.org/gow/db"
	"oldcode.org/gow/server"
)

func main() {
	//	for i, a := range os.Args[1:] {
	//		fmt.Printf("%d:%s ", i, a)
	//	}

	if len(os.Args) >= 2 {
		switch arg1 := os.Args[1]; arg1 {
		case "web":
			web()
		case "db":
			dbstuff()
		default:
			usage()
		}
	} else {
		usage()
		web()
	}

}

func web() {
	cd()
	server.Serve()
}

func cd() {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	ss := []string{gopath, "src", "oldcode.org", "gow"}
	dir := strings.Join(ss, "/")
	fmt.Printf("changing directory to:%s\n", dir)
	if err := os.Chdir(dir); err != nil {
		panic(err)
	}
}

func dbstuff() {
	db.Drivers()
	db.InitDB()
	db.TestDB()
}

func usage() {
	fmt.Println()
	fmt.Printf("gow [web|db]\n")
	fmt.Println()
}

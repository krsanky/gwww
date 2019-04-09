package main

import (
	"fmt"
	"go/build"
	"os"
	"strings"

	"oldcode.org/gow/db"
	"oldcode.org/gow/routes"
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
	routes.Serve()
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
	db.TestGormSql()
	db.Drivers()
	//db.TestSql()
	artists, err := db.GetRawArtists()
	if err != nil {
		panic(err)
	}
	for _, a := range artists {
		fmt.Printf("a:%s\n", a)
	}
}

func usage() {
	fmt.Println()
	fmt.Printf("gow [web|db|cd]\n")
	fmt.Println()
}

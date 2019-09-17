package main

import (
	"fmt"
	"os"

	"oldcode.org/home/wise/repo/go/oldcode.org/gow/db"
	"oldcode.org/home/wise/repo/go/oldcode.org/gow/server"
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
		case "tmpl":
			TmplTest()
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
	d, _ := os.Getwd()
	fmt.Printf("cur dir:%s\n", d)
}

func dbstuff() {
	db.Drivers()
	db.InitDB()
	db.TestDB()
}

func usage() {
	fmt.Println()
	fmt.Printf("gow [web|db|tmpl]\n")
	fmt.Println()
}

package main

import (
	"fmt"
	"os"

	"oldcode.org/home/wise/repo/go/gow/db"
	"oldcode.org/home/wise/repo/go/gow/model"
	"oldcode.org/home/wise/repo/go/gow/server"
)

func main() {
	//	for i, a := range os.Args[1:] {
	//		fmt.Printf("%d:%s ", i, a)
	//	}

	if len(os.Args) == 2 {
		fmt.Printf("2 arg\n")
		usage()
		return
	}
	// 1st arg is cmd, 2nd is settings.toml file
	if len(os.Args) > 2 {
		switch arg1 := os.Args[1]; arg1 {
		case "web":
			settings := os.Args[2]
			server.Serve(settings)
		case "db":
			dbstuff()
		case "tmpl":
			TmplTest()
		default:
			usage()
		}
	} else {
		usage()
	}
}

func dbstuff() {
	db.Drivers()
	db.InitDB()
	db.TestDB()
	as, err := model.GetRawArtists()
	if err != nil {
		panic(err)
	}
	for _, a := range as {
		fmt.Printf("a:%s\n", a)
	}
}

func usage() {
	fmt.Println()
	fmt.Printf("gow [web|db|tmpl]\n")
	fmt.Println()
}

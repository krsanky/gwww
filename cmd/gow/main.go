package main

import (
	"fmt"
	"os"

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
			routes.Serve()
		case "db":
			dbstuff()
		default:
			usage()
		}
	} else {
		usage()
		routes.Serve()
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
	fmt.Printf("gow [web|db]\n")
	fmt.Println()
}

package main

import (
	"fmt"
	"os"

	"oldcode.org/gow/db"
	"oldcode.org/gow/web"
)

func main() {
	fmt.Printf("gow\n")
	for i, a := range os.Args[1:] {
		fmt.Printf("%d:%s ", i, a)
	}
	fmt.Println()

	if len(os.Args) >= 2 {
		switch arg1 := os.Args[1]; arg1 {
		case "web":
			web.Serve()
		case "db":
			db.Main()
			db.TestBeets()
		default:
			usage()
		}
	} else {
		usage()
	}

}

func usage() {
	fmt.Println()
	fmt.Println(`gow [web|db]`)
	fmt.Println()
}

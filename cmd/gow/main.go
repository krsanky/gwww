package main

import (
	"fmt"
	"os"

	"oldcode.org/gow/db"
)

func main() {
	fmt.Printf("gow\n")
	for i, a := range os.Args[1:] {
		fmt.Printf("%d:%s ", i, a)
	}
	fmt.Println()

	//	if len(os.Args) >= 2 {
	//		switch arg1 := os.Args[1]; arg1 {
	//		case "webserver":
	//			runWebserver()
	//		case "importurt":
	//			importUrt()
	//		default:
	//			usage()
	//		}
	//	} else {
	//		usage()
	//	}

	db.Main()
}

func usage() {
	fmt.Println()
	fmt.Println(`webplay [webserver|importurt]`)
	fmt.Println()
}

package main

import (
	"fmt"

	"oldcode.org/home/wise/repo/go/oldcode.org/gow/db"
	"oldcode.org/home/wise/repo/go/oldcode.org/gow/model"
	"oldcode.org/home/wise/repo/go/oldcode.org/gow/settings"
)

func main() {
	//	for i, a := range os.Args[1:] {
	//		fmt.Printf("%d:%s ", i, a)
	//	}
	fmt.Printf("gow_test...\n")

	settings.Init("settings.toml")
	db.InitDB()
	fmt.Printf("--\n")
	db.TestDB()

	as, err := model.GetAllArtists()
	//_, err := model.GetArtists("A")
	if err != nil {
		panic(err)
	}
	for i, a := range as {
		fmt.Printf("%d %s\n", i, a)
	}

}

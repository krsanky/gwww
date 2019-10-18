package main

import (
	"fmt"

	"oldcode.org/home/wise/repo/go/gow/db"
	"oldcode.org/home/wise/repo/go/gow/model"
	"oldcode.org/home/wise/repo/go/gow/phrase"
	"oldcode.org/home/wise/repo/go/gow/settings"
)

func db_artists() {
	//as, err := model.GetAllArtists()
	as, err := model.GetArtists("A")
	if err != nil {
		panic(err)
	}
	for i, a := range as {
		fmt.Printf("%d %s\n", i, a)
	}
}

func test_phrase() {
	fmt.Printf("phrase...\n")
	ps, err := phrase.GetPhrases()
	if err != nil {
		panic(err)
	}
	for _, p := range ps {
		fmt.Printf("%s\n", p.String())
	}
}

func main() {
	//	for i, a := range os.Args[1:] {
	//		fmt.Printf("%d:%s ", i, a)
	//	}
	fmt.Printf("gow_test...\n")

	settings.Init("settings.toml")
	db.InitDB()

	//db_artists()
	test_phrase()
}

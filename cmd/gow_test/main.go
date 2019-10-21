package main

import (
	"fmt"
	"os"

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

func dbstuff() {
	db.Drivers()
	db.TestDB()
	as, err := model.GetRawArtists()
	if err != nil {
		panic(err)
	}
	for _, a := range as {
		fmt.Printf("a:%s\n", a)
	}
}

func main() {
	settings.Init("settings.toml")
	db.InitDB()
	if len(os.Args) > 1 {
		switch arg1 := os.Args[1]; arg1 {
		case "tmpl":
			TmplTest()
		case "db":
			//db_artists()
			dbstuff()
		default:
			test_phrase()
		}
	} else {
		test_phrase()
	}
}

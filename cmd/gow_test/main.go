package main

import (
	"fmt"
	"os"

	"oldcode.org/repo/go/gow/db"
	"oldcode.org/repo/go/gow/model"
	"oldcode.org/repo/go/gow/phrase"
	"oldcode.org/repo/go/gow/settings"
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

func func_nil_string_test(str string) {
	fmt.Printf("--- func_nil_string_test(str string) ---\n")
	fmt.Printf("str:%s\n", str)
}

func nil_string_test() {
	fmt.Printf("call func_nil_string_test(str string) with string:\n")
	func_nil_string_test("this-is-a-string asd asd") 

	fmt.Printf("call func_nil_string_test(str string) with nil:\n")
	//func_nil_string_test(nil)
	fmt.Printf("won't compile :)\n")
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
		case "nil-string":
			nil_string_test()
		case "help":
			fmt.Printf("\ngow_test <tmpl|db|nil-string|help>\n\n")
		default:
			test_phrase()
		}
	} else {
		test_phrase()
	}
}

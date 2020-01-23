package main

import (
	"fmt"
	"os"

	"github.com/krsanky/gwww/db"
	"github.com/krsanky/gwww/settings"
)

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
	db.Init()
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
			dbstuff()
		}
	} else {
		dbstuff()
	}
}

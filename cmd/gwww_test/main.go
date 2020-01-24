// Copyright (c) 2020 Paul Wisehart paul@oldcode.org
//
// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.

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

package main

import (
	"fmt"
	"os"

	"oldcode.org/repo/go/gow/db"
	"oldcode.org/repo/go/gow/settings"
	"oldcode.org/repo/go/gow/stocks"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("%s <data-file>\n", os.Args[0])
		os.Exit(1)
	}
	fn := os.Args[1]

	settings.Init("settings.toml")

	db.Init()
	stocks.Init(fn)
	defer stocks.Cleanup()
	//stocks.GetColumnNames()
	err := stocks.ProcessNasdaqFile()
	if err != nil {
		panic(err)
	}
}

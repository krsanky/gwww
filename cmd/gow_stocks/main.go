package main

import (
	"fmt"
	"os"

	"oldcode.org/repo/go/gow/stocks"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("%s <data-file>\n", os.Args[0])
		os.Exit(1)
	}
	fn := os.Args[1]
/*
	fmt.Printf("arg0:%s\n", os.Args[0]) // cms name
	fmt.Printf("download this: ftp://ftp.nasdaqtrader.com/symboldirectory/nasdaqlisted.txt\n")
*/

	err := stocks.LoadFromFile(fn)
	if err != nil {
		panic(err)
	}
}

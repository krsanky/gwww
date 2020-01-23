package main

import (
	"fmt"
	"os"

	"github.com/krsanky/gwww/server"
)

func main() {
	if len(os.Args) == 2 {
		fmt.Printf("2 arg\n")
		usage()
		return
	}
	// 1st arg is cmd, 2nd is settings.toml file
	if len(os.Args) > 2 {
		switch arg1 := os.Args[1]; arg1 {
		case "web":
			settings := os.Args[2]
			server.Serve(settings)
		default:
			usage()
		}
	} else {
		usage()
	}
}

func usage() {
	fmt.Println()
	fmt.Println("gwww [web]")
	fmt.Println()
}
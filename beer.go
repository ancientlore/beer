package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	for _, s := range flag.Args() {
		fmt.Print(s, " ")
	}
	fmt.Println("🍺 🍺 🍺 ")
}

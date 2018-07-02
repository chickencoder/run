package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) > 0 {
		fmt.Println(args)
	} else {
		fmt.Println("Interactive Run 🏃")
	}
}

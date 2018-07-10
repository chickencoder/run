package main

import "github.com/chickencoder/runlang/scanner"

func main() {
	scanner := scanner.NewScanner("Hello Run")
	scanner.Next()
}

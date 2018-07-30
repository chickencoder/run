package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/chickencoder/run/vm"
)

func main() {
	main := flag.Int("main", 0, "Main entry point for program")
	trace := flag.Bool("trace", false, "Trace the program execution")
	asmFile := flag.String("asm", "", "Execute a Run Assembly Program")
	size := flag.Int("stacksize", 1024, "Fixed size of execution stack")
	flag.Parse()

	// sum:
	//		load x
	//		load y
	//		add
	//		ret
	//
	// 		const 10
	//		const 2
	//		store x
	//		store y
	//		call sum, 2
	//		halt
	// x is at address 0x00
	// y is at address 0x01

	path, err := filepath.Abs(*asmFile)

	dat, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("FileError: couldn't open file ", path)
	}

	program := vm.Assemble(string(dat))
	runner := vm.NewRunner(program, *size, *main, *trace)
	runner.Run()
}

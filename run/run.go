// Copyright 2018 Jesse Sibley. All rights reserved.
// This document, along with the entire source code for the
// Run project is licensed under the BSD-3-Clause License.
// See LICENSE file for more details.

package run

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/chickencoder/run/run/lexer"
)

// Exec interprets a string then executes
func Exec(source string) {
	tok := lexer.NewToken(lexer.StringToken, "Hello World", 0, 10)
	fmt.Println(tok)
	// Error(-1, "Run doesn't exist yet")
}

// ExecFile reads then executes a run script file
func ExecFile(filename string) {
	path, err := filepath.Abs(filename)
	if err != nil {
		fmt.Println(err)
	}

	b, err := ioutil.ReadFile(string(path))
	if err != nil {
		fmt.Println(err)
	}

	Exec(string(b))
}

// ExecRepl initialising a new REPL instance
func ExecRepl() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		Exec(text)
		fmt.Println(text)
	}
}

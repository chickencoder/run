// Copyright 2018 Jesse Sibley. All rights reserved.
// This document, along with the entire source code for the
// Run project is licensed under the BSD-3-Clause License.
// See LICENSE file for more details.

package main

import (
	"flag"
	"fmt"

	"github.com/chickencoder/run/run"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) > 1 {
		fmt.Println("Usage: run [script]")
	} else if len(args) == 1 {
		run.ExecFile(args[0])
	} else {
		fmt.Println("Interactive Run 🏃")
		run.ExecRepl()
	}
}

// Copyright 2018 Jesse Sibley. All rights reserved.
// This document, along with the entire source code for the
// Run project is licensed under the BSD-3-Clause License.
// See LICENSE file for more details.

package run

import "fmt"

// Error reports an error within line of source code
func Error(line int, message string) {
	Report(line, "", message)
}

// Report writes format source-based error messages
func Report(line int, where string, message string) {
	fmt.Printf("\nError: %s\n", message)
	if line != -1 {
		fmt.Printf("%d | %s\n", line, where)
	}
	fmt.Println()
}

// Copyright 2018 Jesse Sibley. All rights reserved.
// This document, along with the entire source code for the
// Run project is licensed under the BSD-3-Clause License.
// See LICENSE file for more details.

package lexer

type TokenType int

const (
	StringToken TokenType = iota
	IdentToken
)

type Token struct {
	Type  TokenType
	Value string
	Start int
	End   int
}

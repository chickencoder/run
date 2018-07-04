// Copyright 2018 Jesse Sibley. All rights reserved.
// This document, along with the entire source code for the
// Run project is licensed under the BSD-3-Clause License.
// See LICENSE file for more details.

package lexer

import "fmt"

// TokenType is used to represent the nature of a token
type TokenType int

// Token Constants are TokenType
const (
	StringToken TokenType = iota
	NumberToken
	BooleanToken

	IdentifierToken
	KeywordToken

	LeftParenToken
	RightParenToken

	LeftBraceToken
	RightBraceToken

	LeftSquareToken
	RightSquareToken

	CommaToken
	DotToken
	MinusToken
	PlusToken
	SlashToken
	StartToken

	BangToken
	BangEqualToken
	EqualToken
	EqualEqualToken
	GreaterToken
	GreaterEqualToken
	LessToken
	LessEqualToken

	EOF
)

// Token represents a lexeme within the lexer
type Token struct {
	Type  TokenType
	Value string
	Start int
	End   int
}

// NewToken returns a new token
func NewToken(typ TokenType, val string, start int, end int) Token {
	return Token{
		typ,
		val,
		start,
		end,
	}
}

func (t Token) String() string {
	return fmt.Sprintf("[%d:%d] %q", t.Start, t.End, t.Value)
}

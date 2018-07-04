// Copyright 2018 Jesse Sibley. All rights reserved.
// This document, along with the entire source code for the
// Run project is licensed under the BSD-3-Clause License.
// See LICENSE file for more details.

package lexer

type runeNode struct {
	r    rune
	next *runeNode
}

type runeStack struct {
	start *runeNode
}

func newRuneStack() runeStack {
	return runeStack{}
}

func (s *runeStack) push(r rune) {
	node := &runeNode{r: r}
	if s.start == nil {
		s.start = node
	} else {
		node.next = s.start
		s.start = node
	}
}

func (s *runeStack) pop() rune {
	if s.start == nil {
		return EOFRune
	}

	n := s.start
	s.start = n.next
	return n.r
}

func (s *runeStack) clear() {
	s.start = nil
}

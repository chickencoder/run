package scanner

import (
	"fmt"
	"unicode"
)

const eof = ""

type Scanner struct {
	start  int
	cursor int
	line   int
	source string
	Tokens chan Token
}

func NewScanner(source string) *Scanner {
	return &Scanner{
		source: source,
		Tokens: make(chan Token, 2),
	}
}

func (l *Scanner) Peek() string {
	if l.cursor < len([]rune(l.source)) {
		return string([]rune(l.source)[l.cursor+1])
	} else {
		return eof
	}
}

func (l *Scanner) Current() string {
	if l.cursor < len([]rune(l.source)) {
		return string([]rune(l.source)[l.cursor])
	} else {
		return eof
	}
}

func (l *Scanner) Emit(typ TokenType) {
	val := l.source[l.start:l.cursor]
	token := Token{
		Type:  typ,
		Value: val,
	}

	l.start = l.cursor
	l.Tokens <- token
}

func (l *Scanner) Step() {
	l.cursor++
}

func (l *Scanner) Error(msg string) {
	fmt.Println("Scanning Error Occured")
	fmt.Printf("Line %d, at %d: %s\n", l.line, l.cursor, msg)
}

func (l *Scanner) Next() {
	if l.cursor < len([]rune(l.source)) {
		if isWhiteSpace(l.Current()) {
			for isWhiteSpace(l.Peek()) {
				l.Step()
			}
		}

		if isDigit(l.Current()) {
			for isDigit(l.Peek()) {
				fmt.Println(l.Current())
				l.Step()
			}
			l.Emit(NumberToken)
		}

	} else {
		l.Emit(EOF)
	}
}

func isDigit(s string) bool {
	r := []rune(s)
	return '0' <= r[0] && r[0] <= '9'
}

func isIdentifier(s string) bool {
	first := true
	for _, r := range s {
		if unicode.IsDigit(r) {
			if first {
				return false
			}
		} else if r != '_' && !unicode.IsLetter(r) {
			return false
		}
		first = false
	}

	return true
}

func isWhiteSpace(s string) bool {
	return s == "\r" || s == "\n" || s == "\t" || s == " "
}

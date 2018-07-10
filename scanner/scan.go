package scanner

import "fmt"

type Scanner struct {
	start  int
	cursor int
	line   int
	source string
	tokens chan Token
}

func NewScanner(source string) *Scanner {
	return &Scanner{
		source: source,
		tokens: make(chan Token, 2),
	}
}

func (l *Scanner) Peek() string {
	return string([]rune(l.source[l.cursor+1]))
}

func (l *Scanner) Emit(typ TokenType) {
	val := l.source[l.start:l.cursor]
	token := Token{
		Type:  typ,
		Value: val,
	}

	l.tokens <- token
}

func (l *Scanner) Next() {
	fmt.Println(l.source)
}

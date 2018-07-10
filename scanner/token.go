package scanner

type TokenType int

const (
	// Literals
	StringToken TokenType = iota
	NumberToken
	BooleanToken
	IdentiferToken

	// Punctuation
	LeftParenToken
	RightParenToken
	LeftBraceToken
	RightBraceToken

	PlusToken
	MinusToken
	StarToken
	SlashToken
	CommaToken
	DotToken

	// Operators
	BangToken
	BangEqualToken
	EqualToken
	EqualEqualToken
	GreaterToken
	GreaterEqualToken
	LessToken
	LessEqualToken

	// Keywords
	AndToken
	OrToken
	IfToken
	ElseToken
	ForToken
	OfToken
	NilToken
	LetToken
	SetToken
	FunToken
	ReturnToken
	ImportToken
	EntityToken

	EOF
)

type Token struct {
	Type  TokenType
	Value string
}

func NewToken(typ TokenType, val string) *Token {
	return &Token{
		Type:  typ,
		Value: val,
	}
}

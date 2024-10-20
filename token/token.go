package token

type TokenType int

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ERROR TokenType = iota
	EOF
	ATOM
	CONST
	NUMBER
	LPAREN
	RPAREN
	PLUS
	MINUS
	DOT
	CHAR
	QUOTE
	NEWLINE
)

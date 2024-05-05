package token

type TokenType string

// Token struct holds detail of a token, i.e,
// the type of a token and the value it represents
type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identitifier + literal
	IDENT = "IDENT" // add, foobar, x, y ...
	INT   = "INT"   // 12345

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimeters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

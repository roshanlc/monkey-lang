package lexer

import "github.com/roshanlc/interpreter-in-go/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

// New returns a lexer
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// readChar reads character from input string
func (l *Lexer) readChar() {
	// check if the read position is greater than input size
	if l.readPosition >= len(l.input) {
		l.ch = 0 // In ASCII, 0 represents "NUL"
	} else {
		l.ch = l.input[l.readPosition]
	}
	// set current position to read position
	l.position = l.readPosition
	// increment the read position
	l.readPosition += 1
}

// NextToken returns the next token from input string
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	// check the value of current character
	switch l.ch {
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

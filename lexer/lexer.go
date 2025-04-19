package lexer

import "Brioche/token"

// This struct is basically defining ecerything needed to read through input string one character at a time
type Lexer struct {
	input string
	// position is always at curr charcter while readPosition is the next character. Basically a 2 pointer approach
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	// Basically a shorter version of l := new(Lexer) and then l.input = input
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		// 0 is ASCII code for null
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
		l.readPosition += 1
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	switch l.ch {
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
	case '+':
		tok = newToken(token.PLUS, l.ch)
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

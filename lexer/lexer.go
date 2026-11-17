package lexer

import (
	"github.com/shantanuparte/vex/token"
)

type Lexer struct {
	input   string
	curr    int  //current position
	readPos int  //next to currect position
	ch      byte //current char in examination
}

// creates new lexer
func New(input string) *Lexer {
	l := Lexer{input: input}
	l.readChar()
	return &l
}

// read the nexer char
func (l *Lexer) readChar() {

	// idk really know why i need the curr in this lol

	if l.readPos >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPos]
	}

	l.curr = l.readPos

	l.readPos++

}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)

	case '+':
		tok = newToken(token.PLUS, l.ch)

	case '(':
		tok = newToken(token.LBRACE, l.ch)

	case ')':
		tok = newToken(token.RBRACE, l.ch)

	case '{':
		tok = newToken(token.LCURLBRACE, l.ch)
	case '}':
		tok = newToken(token.RCURLBRACE, l.ch)
	case ',':
		tok = newToken(token.SEMICOLON, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)

	}

	l.readChar()
	return tok
}

func newToken(tokType token.TokenType, ch byte) token.Token {

	return token.Token{Type: tokType, Litral: string(ch)}

}



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

func (l *Lexer) skipComments() {
	for l.ch != '\n' && l.ch != 0 {
		l.readChar()
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	for {
		l.skipWhitespace()
		if l.ch == '/' && l.peekFurther() == '/' {
			l.skipComments()
		} else {
			break
		}
	}

	switch l.ch {
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
		tok = newToken(token.COMMA, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '<':
		tok = newToken(token.LF, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case '=':
		if l.peekFurther() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.EQUALTO, Litral: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '!':
		if l.peekFurther() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.NOTEUQAL, Litral: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.EXCLAMATION, l.ch)
		}

	case 0:
		tok.Litral = ""
		tok.Type = token.EOFILE

	default:
		if isLetter(l.ch) {
			//need to serprate type n litral->(variable name)
			tok.Litral = l.readIndentifier()
			tok.Type = token.LookIdentfiers(tok.Litral)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Litral = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readNumber() string {
	position := l.curr
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.curr]
}

func newToken(tokType token.TokenType, ch byte) token.Token {

	return token.Token{Type: tokType, Litral: string(ch)}

}

func isLetter(ch byte) bool {
	//user rune->int32 as they are unicode n we can direcly compare then with byte->uint8
	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch == '_'
}

func (l *Lexer) readIndentifier() string {

	//I guess curr is used cause in return slice omit first ittration.

	prevPos := l.curr

	for isLetter(l.ch) || isDigit(l.ch) {
		l.readChar()
	}

	return l.input[prevPos:l.curr]

}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' || l.ch == '\n' {
		l.readChar()
	}
}

func (l *Lexer) peekFurther() byte {
	//check if program is complete-> lol i would have never though abt this endcase
	if l.readPos >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPos] //next character peek
	}
}

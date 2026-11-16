package lexer

import (
	"testing"

	"github.com/shantanuparte/vex/token"
)

func TestNextToken(t *testing.T) {

	test_str := `=+(){},;` // `` for raw strig litrals->without \n\t stuff

	//slice of structs
	tests := []struct {
		extectedType   token.TokenType
		extectedLitral string
	}{

		//arange in order of input lol
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LBRACE, "("},
		{token.RBRACE, ")"},
		{token.LCURLBRACE, "{"},
		{token.RCURLBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOFILE, ""},
	}

	l := New(test_str) //create a new lexser

	// i-> row,
	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.extectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.extectedType, tok.Type)
		}

		if tok.Litral != tt.extectedLitral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.extectedLitral, tok.Literal)
		}

	}

}

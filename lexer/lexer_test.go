package lexer

import (
	"testing"

	"github.com/shantanuparte/vex/token"
)

func TestNextToken(t *testing.T) {

	test_str := `val x = 64;
	val y = 10;
	val sum = fn(x , y){
		x + y;
	}
	val result = sum(x,y);
	!-/*5;
	5 < 10 > 5;
	if (5 < 10) {
	return true;
	} else {
	return false;
		}
	10 == 10;
	10 != 9;
	` 
	
	// `` for raw strig litrals->without \n\t stuff

	//slice of structs
	tests := []struct {
	extectedType   token.TokenType
	extectedLitral string
}{
	// val x = 64;
	{token.VAL, "val"},
	{token.IDENTIFIER, "x"},
	{token.ASSIGN, "="},
	{token.INT, "64"},
	{token.SEMICOLON, ";"},

	// val y = 10;
	{token.VAL, "val"},
	{token.IDENTIFIER, "y"},
	{token.ASSIGN, "="},
	{token.INT, "10"},
	{token.SEMICOLON, ";"},

	// val sum = fn(x, y) { x + y; }
	{token.VAL, "val"},
	{token.IDENTIFIER, "sum"},
	{token.ASSIGN, "="},
	{token.FUNCTION, "fn"},
	{token.LBRACE, "("},
	{token.IDENTIFIER, "x"},
	{token.COMMA, ","},
	{token.IDENTIFIER, "y"},
	{token.RBRACE, ")"},
	{token.LCURLBRACE, "{"},
	{token.IDENTIFIER, "x"},
	{token.PLUS, "+"},
	{token.IDENTIFIER, "y"},
	{token.SEMICOLON, ";"},
	{token.RCURLBRACE, "}"},

	// val result = sum(x, y);
	{token.VAL, "val"},
	{token.IDENTIFIER, "result"},
	{token.ASSIGN, "="},
	{token.IDENTIFIER, "sum"},
	{token.LBRACE, "("},
	{token.IDENTIFIER, "x"},
	{token.COMMA, ","},
	{token.IDENTIFIER, "y"},
	{token.RBRACE, ")"},
	{token.SEMICOLON, ";"},

	// !-/*5;
	{token.EXCLAMATION, "!"},
	{token.MINUS, "-"},
	{token.SLASH, "/"},
	{token.ILLEGAL, "*"}, // Your lexer doesn't handle * yet
	{token.INT, "5"},
	{token.SEMICOLON, ";"},

	// 5 < 10 > 5;
	{token.INT, "5"},
	{token.LF, "<"},
	{token.INT, "10"},
	{token.GT, ">"},
	{token.INT, "5"},
	{token.SEMICOLON, ";"},

	// if (5 < 10) {
	{token.IF, "if"},
	{token.LBRACE, "("},
	{token.INT, "5"},
	{token.LF, "<"},
	{token.INT, "10"},
	{token.RBRACE, ")"},
	{token.LCURLBRACE, "{"},

	// return true;
	{token.RETURN, "return"},
	{token.TRUE, "true"},
	{token.SEMICOLON, ";"},

	// } else {
	{token.RCURLBRACE, "}"},
	{token.ELSE, "else"},
	{token.LCURLBRACE, "{"},

	// return false;
	{token.RETURN, "return"},
	{token.FALSE, "false"},
	{token.SEMICOLON, ";"},

	// }
	{token.RCURLBRACE, "}"},

	// 10 == 10;
	{token.INT, "10"},
	{token.EQUALTO, "=="},
	{token.INT, "10"},
	{token.SEMICOLON, ";"},

	// 10 != 9;
	{token.INT, "10"},
	{token.NOTEUQAL, "!="},
	{token.INT, "9"},
	{token.SEMICOLON, ";"},

	// EOF
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
				i, tt.extectedLitral, tok.Litral)
		}

	}

}

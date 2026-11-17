package token

type TokenType string

type Token struct{
	Type TokenType
	Litral string
}

//tokens 

const(
	//special tokens
	ILLEGAL = "ILLEGAL" //if lexer doesn't understnad the keyword
	EOFILE = "EOFILE"

	IDENTIFIER = "IDENT"
	INT = "INT"

	//operators
	ASSIGN = "="
	PLUS = "+"	
	MINUS = "-"

	//signs
	COMMA = ","
	SEMICOLON = ";" //using for learning will cancel it after this 

	LBRACE = "("
	RBRACE = ")"
	LCURLBRACE = "{"
	RCURLBRACE = "}"

	FUNCTION = "FUNCTION"
	LET = "LET"

)



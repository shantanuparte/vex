package lexer

type Lexer struct{
	input string 
	curr int	//current position 
	readPos int	//next to currect position
	ch byte	//current char in examination
}


func New(input string) *Lexer {
	l := Lexer{input: input}
	return &l
}




package main

import (
	"fmt"
	"os"

	"github.com/shantanuparte/vex/lexer"
	"github.com/shantanuparte/vex/repl"
	"github.com/shantanuparte/vex/token"
)

func main() {
	if len(os.Args) > 1 {
		//use vex file
		run(os.Args[1])

	} else {
		repl.Start(os.Stdin, os.Stdout)
	}
}

func run(path string) {

	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading file: ", err)
		os.Exit(1)
	}

	idk(string(data))

}

func idk(src string) {
	l := lexer.New(src)
	for tok := l.NextToken(); tok.Type != token.EOFILE; tok = l.NextToken() {
		fmt.Printf("%+v\n", tok)
	}

}

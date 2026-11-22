package main

import (
	"fmt"
	"os"

	"github.com/shantanuparte/vex/repl"
)

func main() {

	fmt.Printf("Vex programming language: \n")
	repl.Start(os.Stdin, os.Stdout)

}

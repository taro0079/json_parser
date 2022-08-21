package main

import (
	"fmt"
	"json_parser/json_lexer"
	"json_parser/json_parser"
)

func main() {
	input := `{"test":["ttt","123"]}`
	l := json_lexer.New(input)
	p := json_parser.New(l)
	fmt.Println(p.ParseProgram())
}

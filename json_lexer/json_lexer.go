package json_lexer

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

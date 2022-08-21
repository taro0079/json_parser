package json_lexer

import (
	"json_parser/json_token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0

	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() json_token.Token {
	var tok json_token.Token
	l.skipWhitespace()

	switch l.ch {
	case '{':
		tok = newToken(json_token.LBRACE, l.ch)
	case '}':
		tok = newToken(json_token.RBRACE, l.ch)
	case '[':
		tok = newToken(json_token.LBRAKET, l.ch)
	case ']':
		tok = newToken(json_token.RBRAKET, l.ch)
	case ':':
		tok = newToken(json_token.COLON, l.ch)
	case ',':
		tok = newToken(json_token.COMMA, l.ch)
	case '"':
		tok = newToken(json_token.DOUBLEQUOTE, l.ch)

	case 0:
		tok.Literal = ""
		tok.Type = json_token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = json_token.IDENT
			return tok
		}
	}
	l.readChar()
	return tok

}

func newToken(tokenType json_token.Type, ch byte) json_token.Token {
	return json_token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

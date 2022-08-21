package json_parser

import (
	"fmt"
	"json_parser/json_lexer"
	"json_parser/json_token"
)

type Parser struct {
	l         *json_lexer.Lexer
	curToken  json_token.Token
	peekToken json_token.Token
}

func New(l *json_lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken()
	p.nextToken()
	return p

}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() []string {

	var tokenArray []string
	for p.curToken.Type != json_token.EOF {
		fmt.Println(p.curToken)
		tokenArray = append(tokenArray, p.curToken.Literal)
		p.nextToken()
	}
	return tokenArray
}

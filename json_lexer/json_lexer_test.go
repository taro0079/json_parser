package json_lexer

import (
	"json_parser/json_token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `{"test":["ttt":"123"]}`
	tests := []struct {
		expectedType    json_token.Type
		expectedLiteral string
	}{
		{json_token.LBRACE, "{"},
		{json_token.DOUBLEQUOTE, "\""},
		{json_token.IDENT, "test"},
		{json_token.DOUBLEQUOTE, "\""},
		{json_token.COLON, ":"},
		{json_token.LBRAKET, "["},
		{json_token.DOUBLEQUOTE, "\""},
		{json_token.IDENT, "ttt"},
		{json_token.DOUBLEQUOTE, "\""},
		{json_token.COLON, ":"},
		{json_token.DOUBLEQUOTE, "\""},
		{json_token.INT, "123"},
		{json_token.DOUBLEQUOTE, "\""},
		{json_token.RBRAKET, "]"},
		{json_token.RBRACE, "}"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokenType wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}

}

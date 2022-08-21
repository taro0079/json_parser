package json_token

type Type string

type Token struct {
	Type    Type
	Literal string
}

const (
	EOF   = "EOF"
	IDENT = "IDENT"
	COLON = ":"
	COMMA = ","

	LPAREN      = "("
	RPAREN      = ")"
	LBRAKET     = "["
	RBRAKET     = "]"
	LBRACE      = "{"
	RBRACE      = "}"
	DOUBLEQUOTE = "\""
)

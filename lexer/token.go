package lexer

type TokenKind string

const (
	Number        TokenKind = "Number"
	Identifier    TokenKind = "Identifier"
	Keyword       TokenKind = "Keyword"
	Operator      TokenKind = "Operator"
	Delimiter     TokenKind = "Delimiter"
	StringLiteral TokenKind = "StringLiteral"
)

type Token struct {
	Kind  TokenKind
	Value string
}

package tokenizers

import "github.com/dywoq/dywoqlang/lexer/token"

type Comma struct{}

func (Comma) Tokenize(char string) token.Token {
	if char == "," {
		return token.Token{Kind: token.Comma, Value: char}
	}
	return token.Token{Kind: token.Unknown, Value: char}
}

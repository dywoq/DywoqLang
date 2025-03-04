package tokenizers

import (
	"strings"

	"github.com/dywoq/dywoqlang/lexer/token"
)

type Paren struct{}

func (Paren) Tokenize(char string) token.Token {
	if strings.HasPrefix(char, "(") {
		return token.Token{Kind: token.LeftParen, Value: "("}
	}
	if strings.HasPrefix(char, ")") {
		return token.Token{Kind: token.RightParen, Value: ")"}
	}
	return token.Token{Kind: token.Unknown, Value: char}
}
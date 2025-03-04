package tokenizers

import (
	"strings"

	"github.com/dywoq/dywoqlang/lexer/token"
)

type Comma struct{}

func (Comma) Tokenize(char string) token.Token {
	if strings.HasPrefix(char, ",") {
		return token.Token{Kind: token.Comma, Value: char}
	}
	return token.Token{Kind: token.Unknown, Value: char}
}

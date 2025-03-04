package tokenizers

import (
	"regexp"

	"github.com/dywoq/dywoqlang/lexer/token"
)

type Identifier struct{}

func (i Identifier) Tokenize(char string) token.Token {
	re := regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*$`)
	if re.MatchString(char) {
		return token.Token{Kind: token.Identifier, Value: char}
	}

	return token.Token{Kind: token.Unknown, Value: char}
}

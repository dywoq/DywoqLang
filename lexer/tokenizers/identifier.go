package tokenizers

import (
	"regexp"

	"github.com/dywoq/dywoqlang/lexer/token"
)

type Identifier struct{}

func (Identifier) Tokenize(char string) token.Token {
	re := regexp.MustCompile(`^([a-zA-Z_][a-zA-Z0-9_]*)`)
	matches := re.FindStringSubmatch(char)
	if len(matches) > 1 {
		return token.Token{Kind: token.Identifier, Value: matches[1]}
	}

	return token.Token{Kind: token.Unknown, Value: char}
}

package tokenizers

import (
	"regexp"

	"github.com/dywoq/dywoqlang/lexer/token"
)

type Number struct{}

func (n Number) Tokenize(char string) token.Token {
	re := regexp.MustCompile(`^\d+$`)
	if re.MatchString(char) {
		return token.Token{Kind: token.Number, Value: char}
	}

	return token.Token{Kind: token.Unknown, Value: char}
}

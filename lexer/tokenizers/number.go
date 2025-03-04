package tokenizers

import (
	"regexp"

	"github.com/dywoq/dywoqlang/lexer/token"
)

type Number struct{}

func (Number) Tokenize(char string) token.Token {
	re := regexp.MustCompile(`^(\d+)`)
	matches := re.FindStringSubmatch(char)
	if len(matches) > 1 {
		return token.Token{Kind: token.Number, Value: matches[1]}
	}
	return token.Token{Kind: token.Unknown, Value: char}
}

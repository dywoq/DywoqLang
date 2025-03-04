package tokenizers

import "github.com/dywoq/dywoqlang/lexer/token"

type EOF struct{}

func (EOF) Tokenize(char string) token.Token {
	if char == "" {
		return token.Token{Kind: token.EOF, Value: char}
	}
	return token.Token{Kind: token.Unknown, Value: char}
}

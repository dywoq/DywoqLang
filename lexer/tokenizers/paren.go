package tokenizers

import "github.com/dywoq/dywoqlang/lexer/token"

type Paren struct{}

func (Paren) Tokenize(char string) token.Token {
	if char == "(" {
		return token.Token{Kind: token.LeftParen, Value: char}
	}

	if char == ")" {
		return token.Token{Kind: token.RightParen, Value: char}
	}

	return token.Token{Kind: token.Unknown, Value: char}
}

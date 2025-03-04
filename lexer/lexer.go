package lexer

import (
	"strings"

	"github.com/dywoq/dywoqlang/lexer/token"
)

type Lexer struct {
	Tokenizers []Tokenizer
}

func (l Lexer) Tokenize(input string) []token.Token {
	input = strings.TrimSpace(input)
	if input == "" {
		return nil
	}

	var tokens []token.Token
	remaining := input

	for len(remaining) > 0 {
		if tok := l.findToken(remaining); tok.Kind != token.Unknown {
			tokens = append(tokens, tok)
			remaining = remaining[len(tok.Value):]
			continue
		}
		remaining = remaining[1:]
	}

	return tokens
}

func (l Lexer) findToken(input string) token.Token {
	for _, tokenizer := range l.Tokenizers {
		tok := tokenizer.Tokenize(input)
		if tok.Kind != token.Unknown {
			return tok
		}
	}
	return token.Token{Kind: token.Unknown}
}
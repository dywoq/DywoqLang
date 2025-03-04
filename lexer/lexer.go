package lexer

import "github.com/dywoq/dywoqlang/lexer/token"

type Lexer struct {
	Tokenizers   []Tokenizer
}

func (l Lexer) Tokenize(input string) []token.Token {
	tokens := []token.Token{}

	for _, r := range input {
		for _, tokenizer := range l.Tokenizers {
			tokens = append(tokens, tokenizer.Tokenize(string(r)))
		}
	}

	return tokens
}

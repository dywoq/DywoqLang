package lexer

import "github.com/dywoq/dywoqlang/lexer/token"

type Tokenizer interface {
	Tokenize(char string) token.Token
}

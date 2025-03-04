package main

import (
	"fmt"

	"github.com/dywoq/dywoqlang/lexer"
	"github.com/dywoq/dywoqlang/lexer/tokenizers"
	"github.com/dywoq/dywoqlang/lexer/tokenizers/arithmetic"
)

func main() {
	lex := lexer.Lexer{
		Tokenizers: []lexer.Tokenizer{
			tokenizers.Comma{},
			tokenizers.EOF{},
			tokenizers.Identifier{},
			tokenizers.Paren{},
			tokenizers.Number{},
			arithmetic.Operator{},
		},
	}

	fmt.Println(lex.Tokenize("(1 + 2) + (3 / 2 + (3 + 31))"))
}

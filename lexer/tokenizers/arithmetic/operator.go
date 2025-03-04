package arithmetic

import (
	"github.com/dywoq/dywoqlang/lexer/token"
	"github.com/dywoq/dywoqlang/lexer/token/arithmetic"
)

type Operator struct{}

func (Operator) Tokenize(char string) token.Token {
	if char == "+" {
		return token.Token{Kind: arithmetic.OperatorPlus, Value: char}
	}

	if char == "-" {
		return token.Token{Kind: arithmetic.OperatorMinus, Value: char}
	}

	if char == "/" {
		return token.Token{Kind: arithmetic.OperatorDivide, Value: char}
	}

	if char == "*" {
		return token.Token{Kind: arithmetic.OperatorMultiply, Value: char}
	}

	return token.Token{Kind: token.Unknown, Value: char}
}

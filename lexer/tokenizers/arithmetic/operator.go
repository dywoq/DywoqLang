package arithmetic

import (
	"strings"

	"github.com/dywoq/dywoqlang/lexer/token"
	"github.com/dywoq/dywoqlang/lexer/token/arithmetic"
)

type Operator struct{}

func (Operator) Tokenize(char string) token.Token {
	if strings.HasPrefix(char, "+") {
		return token.Token{Kind: arithmetic.OperatorPlus, Value: "+"}
	}
	if strings.HasPrefix(char, "-") {
		return token.Token{Kind: arithmetic.OperatorMinus, Value: "-"}
	}
	if strings.HasPrefix(char, "/") {
		return token.Token{Kind: arithmetic.OperatorDivide, Value: "/"}
	}
	if strings.HasPrefix(char, "*") {
		return token.Token{Kind: arithmetic.OperatorMultiply, Value: "*"}
	}
	return token.Token{Kind: token.Unknown, Value: char}
}

package lexer

type Tokenizer interface {
	Tokenize(char string) string
}

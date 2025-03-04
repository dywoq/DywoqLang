package token

type Kind string

const (
	Number      Kind = "Number"
	Identifier  Kind = "Identifier"
	Operator    Kind = "Operator"
	Punctuation Kind = "Punctuation"
	EOF         Kind = "EOF"
	LeftParen   Kind = "LeftParen"
	RightParen  Kind = "RightParen"
	Unknown     Kind = "Unknown"
)

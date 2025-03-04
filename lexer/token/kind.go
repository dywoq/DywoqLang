package token

type Kind string

const (
	Number                     Kind = "Number"
	Identifier                 Kind = "Identifier"
	Punctuation                Kind = "Punctuation"
	EOF                        Kind = "EOF"
	LeftParen                  Kind = "LeftParen"
	RightParen                 Kind = "RightParen"
	Unknown                    Kind = "Unknown"
)

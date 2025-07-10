package tokens

// TokenType represents the general classification of a token.
type TokenType int

const (
	Illegal TokenType = iota
	Keyword
	Identifier
	Literal
	Operator
	Comment
	Whitespace
	Punctuation
	Other
)

package tokens

// TokenType represents the general classification of a token.
type TokenType int

const (
	Illegal    TokenType = iota
	Keyword              // e.g. SELECT, UPDATE
	Identifier           // table or column names
	Number               // numeric literal
	String               // quoted string literal
	Operator             // + - = etc
	Comment
	Whitespace
	Punctuation
	Other
)

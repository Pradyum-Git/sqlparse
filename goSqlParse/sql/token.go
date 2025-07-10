package sql

import "gosqlparse/tokens"

// Token represents a single SQL token.
type Token struct {
	Type  tokens.TokenType
	Value string
}

func (t Token) String() string {
	return t.Value
}

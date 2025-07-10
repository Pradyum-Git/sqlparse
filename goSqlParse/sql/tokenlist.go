package sql

// TokenList represents a list of tokens.
type TokenList struct {
	Tokens []Token
}

// Len returns number of tokens.
func (tl *TokenList) Len() int { return len(tl.Tokens) }

// Append adds a token to the list.
func (tl *TokenList) Append(tok Token) { tl.Tokens = append(tl.Tokens, tok) }

package sql

import (
	"gosqlparse/tokens"
)

// Token represents a lexical token with optional hierarchical info.
type Token struct {
	Type         tokens.TokenType
	Value        string
	Parent       *TokenList
	IsKeyword    bool
	IsGroup      bool
	IsWhitespace bool
	IsNewline    bool
}

// TokenList is a list of tokens that can itself act as a token.
type TokenList struct {
	Token
	Tokens []*Token
}

func NewTokenList(toks []*Token) *TokenList {
	tl := &TokenList{Tokens: toks}
	tl.Type = tokens.Illegal
	tl.IsGroup = true
	for _, t := range tl.Tokens {
		t.Parent = tl
	}
	return tl
}

// Statement represents a full SQL statement.
type Statement struct {
	*TokenList
}

func NewStatement(toks []*Token) *Statement {
	return &Statement{NewTokenList(toks)}
}

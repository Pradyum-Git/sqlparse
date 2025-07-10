package sql

import (
	"gosqlparse/tokens"
)

// Token represents a lexical token with optional hierarchical info.
type Token struct {
	Type         tokens.TokenType
	Value        string
	Parent       *TokenList
	Group        *TokenList
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

func (tl *TokenList) TokenPtr() *Token { return &tl.Token }

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

// Identifier represents a dotted identifier or alias.
type Identifier struct{ *TokenList }

func NewIdentifier(toks []*Token) *Identifier {
	id := &Identifier{NewTokenList(toks)}
	id.Token.Group = id.TokenList
	return id
}

// Function represents a function call with arguments.
type Function struct{ *TokenList }

func NewFunction(toks []*Token) *Function {
	f := &Function{NewTokenList(toks)}
	f.Token.Group = f.TokenList
	return f
}

// Parenthesis represents tokens enclosed by parentheses.
type Parenthesis struct{ *TokenList }

func NewParenthesis(toks []*Token) *Parenthesis {
	p := &Parenthesis{NewTokenList(toks)}
	p.Token.Group = p.TokenList
	return p
}

// ReplaceRange replaces tokens from start to end (inclusive) with the provided token.
func (tl *TokenList) ReplaceRange(start, end int, tok *Token) {
	for i := start; i <= end && i < len(tl.Tokens); i++ {
		tl.Tokens[i].Parent = nil
	}
	tok.Parent = tl
	tl.Tokens = append(tl.Tokens[:start], append([]*Token{tok}, tl.Tokens[end+1:]...)...)
}

func (tl *TokenList) ReplaceRangeWithList(start, end int, grp *TokenList) {
	grp.Token.Parent = tl
	grp.Token.Group = grp
	tl.ReplaceRange(start, end, &grp.Token)
}

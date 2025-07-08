package statementsplitter

import (
	"gosqlparse/sql"
	"gosqlparse/tokens"
)

// Split splits tokens into statements based on semicolons.
func Split(toks []*sql.Token) []*sql.Statement {
	var stmts [][]*sql.Token
	current := []*sql.Token{}
	for _, t := range toks {
		current = append(current, t)
		if t.Type == tokens.Punctuation && t.Value == ";" {
			stmts = append(stmts, current)
			current = []*sql.Token{}
		}
	}
	if len(current) > 0 {
		stmts = append(stmts, current)
	}
	out := make([]*sql.Statement, 0, len(stmts))
	for _, s := range stmts {
		out = append(out, sql.NewStatement(s))
	}
	return out
}

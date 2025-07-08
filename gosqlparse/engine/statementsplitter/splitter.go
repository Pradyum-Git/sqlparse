package statementsplitter

import (

	"strings"



	"gosqlparse/sql"
	"gosqlparse/tokens"
)

// Split splits tokens into statements based on semicolons.
func Split(toks []*sql.Token) []*sql.Statement {
	var stmts [][]*sql.Token
	current := []*sql.Token{}

	level := 0
	for i, t := range toks {
		current = append(current, t)
		if t.Type == tokens.Punctuation {
			if t.Value == "(" {
				level++
			} else if t.Value == ")" && level > 0 {
				level--
			} else if t.Value == ";" && level == 0 {
				stmts = append(stmts, current)
				current = []*sql.Token{}
				continue
			}
		}
		if t.Type == tokens.Keyword && level == 0 {
			if strings.EqualFold(t.Value, "GO") {
				stmts = append(stmts, current[:len(current)-1])
				current = []*sql.Token{}
				continue
			}
		}
		// end-of-tokens
		if i == len(toks)-1 {
			// don't append if already added due to ';' or GO
			if len(current) > 0 {
				stmts = append(stmts, current)
			}
		}
	}
	if len(current) > 0 && len(stmts) == 0 { // handle empty input

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

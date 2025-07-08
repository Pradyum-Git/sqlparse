package grouping

import (
	"gosqlparse/sql"
	"gosqlparse/tokens"
)

// Group is a placeholder grouping step.
func Group(stmt *sql.Statement) {
	groupParenthesis(stmt.TokenList)
}

func groupParenthesis(tl *sql.TokenList) {
	for i := 0; i < len(tl.Tokens); i++ {
		tok := tl.Tokens[i]
		if tok.Type == tokens.Punctuation && tok.Value == "(" {
			// find matching )
			depth := 1
			j := i + 1
			for j < len(tl.Tokens) {
				t := tl.Tokens[j]
				if t.Type == tokens.Punctuation {
					if t.Value == "(" {
						depth++
					} else if t.Value == ")" {
						depth--
						if depth == 0 {
							break
						}
					}
				}
				j++
			}
			if depth == 0 {
				sub := tl.Tokens[i : j+1]
				grp := sql.NewParenthesis(sub)
				tl.ReplaceRangeWithList(i, j, grp.TokenList)
			}
		} else if tok.Group != nil {
			groupParenthesis(tok.Group)
		}
	}
}

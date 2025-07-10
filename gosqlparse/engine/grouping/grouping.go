package grouping

import (
	"gosqlparse/sql"
	"gosqlparse/tokens"
)

// Group is a placeholder grouping step.
func Group(stmt *sql.Statement) {
	groupParenthesis(stmt.TokenList)
	groupFunctions(stmt.TokenList)
	groupIdentifiers(stmt.TokenList)
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

// groupFunctions joins identifier + parenthesis into a function call.
func groupFunctions(tl *sql.TokenList) {
	for i := 0; i < len(tl.Tokens)-1; i++ {
		tok := tl.Tokens[i]
		next := tl.Tokens[i+1]
		if tok.Type == tokens.Identifier && next.Group != nil {
			g := next.Group
			if len(g.Tokens) >= 2 && g.Tokens[0].Value == "(" && g.Tokens[len(g.Tokens)-1].Value == ")" {
				sub := tl.Tokens[i : i+2]
				grp := sql.NewFunction(sub)
				tl.ReplaceRangeWithList(i, i+1, grp.TokenList)
			}
		}
		if tok.Group != nil {
			groupFunctions(tok.Group)
		}
	}
	if len(tl.Tokens) > 0 {
		if last := tl.Tokens[len(tl.Tokens)-1]; last.Group != nil {
			groupFunctions(last.Group)
		}
	}
}

// groupIdentifiers collapses dotted identifiers into one token.
func groupIdentifiers(tl *sql.TokenList) {
	for i := 0; i < len(tl.Tokens); i++ {
		tok := tl.Tokens[i]
		if tok.Type == tokens.Identifier {
			j := i + 1
			for j < len(tl.Tokens)-1 {
				if tl.Tokens[j].Type == tokens.Punctuation && tl.Tokens[j].Value == "." && tl.Tokens[j+1].Type == tokens.Identifier {
					j += 2
				} else {
					break
				}
			}
			if j > i+1 {
				sub := tl.Tokens[i:j]
				grp := sql.NewIdentifier(sub)
				tl.ReplaceRangeWithList(i, j-1, grp.TokenList)
			}
		}
		if tok.Group != nil {
			groupIdentifiers(tok.Group)
		}
	}
}

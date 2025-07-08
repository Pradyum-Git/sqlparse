package filterstack

import (
	"gosqlparse/engine/grouping"
	"gosqlparse/engine/statementsplitter"
	"gosqlparse/lexer"
	"gosqlparse/sql"
)

// FilterStack coordinates tokenization and grouping.
type FilterStack struct {
	doGrouping bool
}

func New() *FilterStack { return &FilterStack{} }

func (fs *FilterStack) EnableGrouping() { fs.doGrouping = true }

// Run parses SQL text and returns statements.
func (fs *FilterStack) Run(sqlText string) []*sql.Statement {
	lex := lexer.New()
	tokens := lex.Tokenize(sqlText)
	stmts := statementsplitter.Split(tokens)
	if fs.doGrouping {
		for _, s := range stmts {
			grouping.Group(s)
		}
	}
	return stmts
}

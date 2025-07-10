package engine

import "gosqlparse/sql"

// Filter defines an interface for token filters.
type Filter interface {
	Process([]sql.Token) []sql.Token
}

// FilterStack coordinates filters and grouping.
type FilterStack struct {
	preprocess []Filter
}

// New returns a new FilterStack.
func New() *FilterStack { return &FilterStack{} }

// Run executes the filter stack on the given SQL string.
func (fs *FilterStack) Run(input string) []*sql.Statement {
	// TODO: tokenizer and grouping will be implemented later.
	return nil
}

package engine


import (
	"gosqlparse/lexer"
	"gosqlparse/sql"
)




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


// Run tokenizes the input and returns a single Statement for now. Future
// implementations will add statement splitting, filters and grouping similar
// to the Python library.
func (fs *FilterStack) Run(input string) []*sql.Statement {
	toks := lexer.Tokenize(input)

	// apply preprocess filters if any
	for _, f := range fs.preprocess {
		toks = f.Process(toks)
	}

	stmt := &sql.Statement{TokenList: sql.TokenList{Tokens: toks}}
	return []*sql.Statement{stmt}

}

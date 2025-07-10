package filters

import "gosqlparse/sql"

// FilterFunc wraps a function to implement engine.Filter.
type FilterFunc func([]sql.Token) []sql.Token

func (f FilterFunc) Process(toks []sql.Token) []sql.Token { return f(toks) }

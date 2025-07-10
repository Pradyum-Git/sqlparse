package lexer

// Keywords lists the SQL keywords recognized by the simple tokenizer.
var Keywords = map[string]struct{}{
	"SELECT": {},
	"FROM":   {},
	"WHERE":  {},
	"INSERT": {},
	"INTO":   {},
	"VALUES": {},
	"UPDATE": {},
	"DELETE": {},
	"LIMIT":  {},
	"JOIN":   {},
	"ON":     {},
	"AND":    {},
	"OR":     {},
	"CASE":   {},
	"WHEN":   {},
	"THEN":   {},
	"END":    {},
	"AS":     {},
	"GROUP":  {},
	"BY":     {},
	"ORDER":  {},
}

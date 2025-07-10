package lexer

import (
	"gosqlparse/sql"
	"gosqlparse/tokens"
)

// Tokenize splits the input SQL string into tokens.
func Tokenize(input string) []sql.Token {
	// TODO: implement regex-based tokenizer.
	return []sql.Token{
		{Type: tokens.Other, Value: input},
	}
}

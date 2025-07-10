package gosqlparse

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	input := "SELECT a, b FROM t WHERE (c = 1);"
	stmts := Parse(input, true)
	for _, stmt := range stmts {
		fmt.Println("Statement:")
		for _, tok := range stmt.Tokens {
			fmt.Printf("%s -> %q\n", tok.Type.String(), tok.Value)
		}
	}
}

package gosqlparse

import (
	"gosqlparse/engine/filterstack"
	"gosqlparse/sql"
)

// Parse returns parsed SQL statements with optional grouping.
func Parse(sqlText string, grouping bool) []*sql.Statement {
	fs := filterstack.New()
	if grouping {
		fs.EnableGrouping()
	}
	return fs.Run(sqlText)
}

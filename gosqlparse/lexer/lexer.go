package lexer

import (
	"regexp"

	"strings"


	"gosqlparse/sql"
	"gosqlparse/tokens"
)

// Lexer implements a simple regex-based scanner.
type Lexer struct {
	patterns []lexEntry

	keywords map[string]struct{}

}

type lexEntry struct {
	regex *regexp.Regexp
	ttype tokens.TokenType
}

func New() *Lexer {
	l := &Lexer{}
	l.patterns = []lexEntry{

		{regexp.MustCompile(`^\r?\n`), tokens.Newline},
		{regexp.MustCompile(`^\s+`), tokens.Whitespace},
		{regexp.MustCompile(`^[(),;]`), tokens.Punctuation},
		{regexp.MustCompile(`^'(?:''|[^'])*'`), tokens.String},
		{regexp.MustCompile(`^(?:<=|>=|<>|!=|=|<|>)`), tokens.Operator},
		{regexp.MustCompile(`^\d+`), tokens.Number},
		{regexp.MustCompile(`^[a-zA-Z_][\w]*`), tokens.Identifier},
	}
	l.keywords = map[string]struct{}{
		"SELECT": {}, "FROM": {}, "WHERE": {}, "INSERT": {},
		"UPDATE": {}, "DELETE": {}, "IN": {}, "VALUES": {},
		"AND": {}, "OR": {}, "NOT": {}, "INTO": {}, "SET": {},
		"CASE": {}, "WHEN": {}, "THEN": {}, "END": {}, "AS": {},
		"ON": {}, "JOIN": {}, "LEFT": {}, "RIGHT": {}, "FULL": {},
		"OUTER": {}, "INNER": {}, "GROUP": {}, "BY": {}, "ORDER": {},
	}

		{regexp.MustCompile(`^\s+`), tokens.Whitespace},
		{regexp.MustCompile(`^[(),;]`), tokens.Punctuation},
		{regexp.MustCompile(`^'(?:''|[^'])*'`), tokens.String},
		{regexp.MustCompile(`^\d+`), tokens.Number},
		{regexp.MustCompile(`^[a-zA-Z_][\w]*`), tokens.Identifier},
	}

	return l
}

// Tokenize converts SQL text into token objects.
func (l *Lexer) Tokenize(sqlText string) []*sql.Token {
	var tokensOut []*sql.Token
	for len(sqlText) > 0 {
		matched := false
		for _, p := range l.patterns {
			if loc := p.regex.FindStringIndex(sqlText); loc != nil && loc[0] == 0 {
				val := sqlText[:loc[1]]
				tok := &sql.Token{
					Type:         p.ttype,
					Value:        val,
					IsWhitespace: p.ttype == tokens.Whitespace,

					IsNewline:    p.ttype == tokens.Newline,
					IsKeyword:    false,
				}
				// check for keyword if identifier
				if p.ttype == tokens.Identifier {
					up := strings.ToUpper(val)
					if _, ok := l.keywords[up]; ok {
						tok.Type = tokens.Keyword
						tok.IsKeyword = true
					}
				}


					IsKeyword:    false,
				}

				tokensOut = append(tokensOut, tok)
				sqlText = sqlText[loc[1]:]
				matched = true
				break
			}
		}
		if !matched {
			// consume one rune as illegal token
			tok := &sql.Token{Type: tokens.Illegal, Value: string(sqlText[0])}
			tokensOut = append(tokensOut, tok)
			sqlText = sqlText[1:]
		}
	}
	return tokensOut
}

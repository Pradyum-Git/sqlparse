package lexer

import (
	"regexp"
	"strings"
	"unicode/utf8"
)


var (
	whitespaceRe  = regexp.MustCompile(`^\s+`)
	commentRe     = regexp.MustCompile(`^(?:--.*?(?:\n|$)|/\*.*?\*/)`)
	numberRe      = regexp.MustCompile(`^[0-9]+(?:\.[0-9]+)?`)
	stringRe      = regexp.MustCompile(`^'(?:''|[^'])*'`)
	identRe       = regexp.MustCompile(`^[a-zA-Z_][\w$]*`)
	operatorRe    = regexp.MustCompile(`^(?:<>|!=|<=|>=|::|[=<>*/+-])`)
	punctuationRe = regexp.MustCompile(`^[(),.;]`)
)

// Tokenize splits the input SQL string into a slice of tokens.
// This is a simplified lexer intended as a starting point for a full port
// of the Python library. It covers basic keywords, identifiers, numbers,
// strings, comments, operators and punctuation.
func Tokenize(input string) []sql.Token {
	tokensOut := []sql.Token{}

	for len(input) > 0 {
		if m := whitespaceRe.FindString(input); m != "" {
			tokensOut = append(tokensOut, sql.Token{Type: tokens.Whitespace, Value: m})
			input = input[len(m):]
			continue
		}

		if m := commentRe.FindString(input); m != "" {
			tokensOut = append(tokensOut, sql.Token{Type: tokens.Comment, Value: m})
			input = input[len(m):]
			continue
		}

		if m := stringRe.FindString(input); m != "" {
			tokensOut = append(tokensOut, sql.Token{Type: tokens.String, Value: m})
			input = input[len(m):]
			continue
		}

		if m := numberRe.FindString(input); m != "" {
			tokensOut = append(tokensOut, sql.Token{Type: tokens.Number, Value: m})
			input = input[len(m):]
			continue
		}

		if m := operatorRe.FindString(input); m != "" {
			tokensOut = append(tokensOut, sql.Token{Type: tokens.Operator, Value: m})
			input = input[len(m):]
			continue
		}

		if m := punctuationRe.FindString(input); m != "" {
			tokensOut = append(tokensOut, sql.Token{Type: tokens.Punctuation, Value: m})
			input = input[len(m):]
			continue
		}

		if m := identRe.FindString(input); m != "" {
			upper := strings.ToUpper(m)
			if _, ok := Keywords[upper]; ok {
				tokensOut = append(tokensOut, sql.Token{Type: tokens.Keyword, Value: m})
			} else {
				tokensOut = append(tokensOut, sql.Token{Type: tokens.Identifier, Value: m})
			}
			input = input[len(m):]
			continue
		}

		// unknown character, emit as Other and advance by one rune
		r, size := utf8.DecodeRuneInString(input)
		tokensOut = append(tokensOut, sql.Token{Type: tokens.Other, Value: string(r)})
		input = input[size:]
	}

	return tokensOut

}

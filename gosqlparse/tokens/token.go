package tokens

// TokenType defines SQL token categories.
// This is a minimal subset for demonstration.
type TokenType int

const (
	Illegal TokenType = iota
	Whitespace
	Keyword
	Identifier
	Operator
	String
	Number
	Punctuation
)

func (t TokenType) String() string {
	switch t {
	case Illegal:
		return "Illegal"
	case Whitespace:
		return "Whitespace"
	case Keyword:
		return "Keyword"
	case Identifier:
		return "Identifier"
	case Operator:
		return "Operator"
	case String:
		return "String"
	case Number:
		return "Number"
	case Punctuation:
		return "Punctuation"
	default:
		return "Unknown"
	}
}

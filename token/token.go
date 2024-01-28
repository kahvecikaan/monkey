package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals

	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1234524

	// Operators

	ASSIGN = "="
	PLUS   = "+"

	// Delimiters

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords

	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// LookupIdent() checks the keywords table to see whether the given identifier is
// in fact a keyword. If it is, it returns the keyword's TokenType constant. If
// it isn't, it assumes we're dealing with a regular identifier and returns the
// token.IDENT constant.

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT // The TokenType for all user-defined identifiers
}

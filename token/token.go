package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

// Strings were chosen for ease of use. It doesn't have the same level of performance as using an int or byte.
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" //add, foobar, x, y, ...
	INT   = "INT"   // 123456

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {
	// Check if the identifier is a keyword
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	// If not, return IDENT
	return IDENT
}

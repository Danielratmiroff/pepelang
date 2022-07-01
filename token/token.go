package token

// Language tokens used by the Lexer

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT  = "IDENT" // add, foobar, x, y, ...
	INT    = "INT"   // 1343456
	STRING = "STRING"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	EQ     = "=="
	NOT_EQ = "!="

	INC = "++"
	DEC = "--"

	// Delimiters
	COMMA     = ","
	COLON     = ":"
	SEMICOLON = ";"

	LBRACKET = "["
	RBRACKET = "]"
	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	VAR      = "VAR"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	// FOR      = "FOR"
)

var keywords = map[string]TokenType{
	"fn":       FUNCTION,
	"var":      VAR,
	"verdad":   TRUE,
	"falso":    FALSE,
	"si":       IF,
	"sino":     ELSE,
	"retornar": RETURN,
	// "por":      FOR,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

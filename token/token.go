package token

// Declaring TokenType as a type of string
type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

// Declaring TokenTypes
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + Literals
	IDENT  = "IDENT"
	INT    = "INT"
	FLOAT  = "FLOAT"
	STRING = "STRING"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	SLASH    = "/"
	ASTERISK = "*"
	EXCITE   = "!"

	LT     = "<"
	GT     = ">"
	EQUAL  = "=="
	NEQUAL = "!="
	PEQUAL = "+="
	MEQUAL = "-="
	TEQUAL = "*="
	DEQUAL = "/="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"
	PERIOD    = "."
	NEWLINE   = "\n"

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	// Keywords
	FUNC   = "FUNC"
	LET    = "LET"
	TRUE   = "TRUE"
	FALSE  = "FALSE"
	IF     = "IF"
	ELSE   = "ELSE"
	RETURN = "RETURN"
)

var keywords = map[string]TokenType{
	"func":   FUNC,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}

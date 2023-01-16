package token

// Declaring TokenType as a type of string
type TokenType int

type Token struct {
	Type    TokenType
	Literal string
}

// Declaring TokenTypes
const (
	ILLEGAL = iota
	EOF

	// Identifiers + Literals
	IDENT
	INT
	STRING
	FLOAT

	// Operators
	ASSIGN
	PLUS
	MINUS
	SLASH
	ASTERISK
	EXCITE

	LT
	GT
	EQUAL
	NEQUAL

	// Delimiters
	COMMA
	SEMICOLON
	COLON
	PERIOD
	NEWLINE

	LPAREN
	RPAREN
	LBRACE
	RBRACE
	LBRACKET
	RBRACKET

	// Keywords
	FUNC
	LET
	TRUE
	FALSE
	IF
	ELSE
	RETURN
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

func lookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}

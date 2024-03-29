package lexer

import "wrench/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhiteSpace()

	switch ch := l.ch; {
	case isLetter(ch):
		tok.Literal = l.readIdentifier()
		tok.Type = token.LookupIdent(tok.Literal)
	case isDigit(ch) || ch == '.':
		tok.Type, tok.Literal = l.readNumber()
	default:
		switch ch {
		case '=':
			if l.peekChar() == '=' {
				ch := l.ch
				l.readChar()
				tok = token.Token{Type: token.EQUAL, Literal: string(ch) + string(l.ch)}
			} else {
				tok = newToken(token.ASSIGN, l.ch)
			}
		case '+':
			if l.peekChar() == '=' {
				ch := l.ch
				l.readChar()
				tok = token.Token{Type: token.PEQUAL, Literal: string(ch) + string(l.ch)}
			} else {
				tok = newToken(token.PLUS, l.ch)
			}
		case '-':
			if l.peekChar() == '=' {
				ch := l.ch
				l.readChar()
				tok = token.Token{Type: token.MEQUAL, Literal: string(ch) + string(l.ch)}
			} else {
				tok = newToken(token.MINUS, l.ch)
			}
		case '*':
			if l.peekChar() == '=' {
				ch := l.ch
				l.readChar()
				tok = token.Token{Type: token.TEQUAL, Literal: string(ch) + string(l.ch)}
			} else {
				tok = newToken(token.ASTERISK, l.ch)
			}
		case '/':
			if l.peekChar() == '=' {
				ch := l.ch
				l.readChar()
				tok = token.Token{Type: token.DEQUAL, Literal: string(ch) + string(l.ch)}
			} else {
				tok = newToken(token.SLASH, l.ch)
			}
		case '!':
			if l.peekChar() == '=' {
				ch := l.ch
				l.readChar()
				tok = token.Token{Type: token.NEQUAL, Literal: string(ch) + string(l.ch)}
			} else {
				tok = newToken(token.EXCITE, l.ch)
			}
		case '<':
			tok = newToken(token.LT, l.ch)
		case '>':
			tok = newToken(token.GT, l.ch)
		case ',':
			tok = newToken(token.COMMA, l.ch)
		case ';':
			tok = newToken(token.SEMICOLON, l.ch)
		case ':':
			tok = newToken(token.COLON, l.ch)
		case '.':
			tok = newToken(token.PERIOD, l.ch)
		case '\n':
			tok = newToken(token.NEWLINE, ' ')
		case '(':
			tok = newToken(token.LPAREN, l.ch)
		case ')':
			tok = newToken(token.RPAREN, l.ch)
		case '{':
			tok = newToken(token.LBRACE, l.ch)
		case '}':
			tok = newToken(token.RBRACE, l.ch)
		case '[':
			tok = newToken(token.LBRACKET, l.ch)
		case ']':
			tok = newToken(token.RBRACKET, l.ch)
		case '"':
			tok.Type = token.STRING
			tok.Literal = l.readString()
		case 0:
			tok.Type = token.EOF
			tok.Literal = ""
		default:
			tok = newToken(token.ILLEGAL, l.ch)
		}
		l.readChar()
	}

	return tok
}

func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' {
		l.readChar()
	}
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}

	return l.input[position:l.position]
}

func (l *Lexer) readNumber() (token.TokenType, string) {
	position := l.position
	tok := token.ILLEGAL

	if isDigit(l.ch) {
		tok = token.INT
		l.scanNumber()
	}

	if l.ch == '.' {
		tok = token.FLOAT
		l.readChar()
		l.scanNumber()
	}

	return token.TokenType(tok), l.input[position:l.position]
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) || isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) scanNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

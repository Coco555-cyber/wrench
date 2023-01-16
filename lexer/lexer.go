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

func (l *Lexer) nextToken() token.Token {
	var tok token.Token

	l.skipWhiteSpace()

	switch l.ch {
	case '=':
	case '+':
	case '-':
	case '/':
	case '*':
	case '!':
	case '<':
	case '>':
	case ',':
	case ';':
	case ':':
	case '.':
	case '\n':
	case '(':
	case ')':
	case '{':
	case '}':
	case '[':
	case ']':
	case '"':
	case 0:
	default:
	}

	l.readChar()
	return tok
}

func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' {
		l.readChar()
	}
}

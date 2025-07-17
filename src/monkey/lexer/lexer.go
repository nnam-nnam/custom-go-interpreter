package lexer

import "monkey/token"

type Lexer struct {
	input        string
	position     int  // current position in input, which should point to current char
	readPosition int  // current reading position in input (after current char)
	ch           byte // current ASCII char under examination
}

// TODO: Handle Unicode and emojis later on?

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // initialize fields in Lexer struct
	return l
}

func (l *Lexer) readChar() {
	// check if we've reached end of input
	if l.readPosition >= len(l.input) {
		l.ch = 0 // set to NUL
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

// helper function to convert char to new Token
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// return next token from input source
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else {
			// encountered unexpected token
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

// reads for an identifier from current position in lexer; returns when non-letter char is encountered
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// check if ASCII character is a letter or an underscore
func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_'
}

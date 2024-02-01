package lexer

import (
	"interpreterInGo/token"
)

type Lexer struct {
	input        string // input string
	position     int    // current position in input (points to current char)
	readPosition int    // current reading position in input (after current char)
	ch           byte   // current char under examination
}

// New returns a new lexer
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// readChar reads the next character in the input and advances our position in the input string
// 获取下一个字符，将会针对这个字符进行判断和处理
func (l *Lexer) readChar() {
	// If we've reached the end of the input, set ch to 0 (ASCII code for "NUL" or "NULL")
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		// Otherwise, set ch to the next character in the input string
		l.ch = l.input[l.readPosition]
	}

	// Advance our position in the input string
	l.position = l.readPosition
	l.readPosition += 1
}

// 查看后一个字符，不会移动指针
func (l *Lexer) peekChar() byte {
	// If we've reached the end of the input, set ch to 0 (ASCII code for "NUL" or "NULL")
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		// Otherwise, set ch to the next character in the input string
		return l.input[l.readPosition]
	}
}

// 跳过空白字符直到当前字符为非空白字符
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// readIdentifier reads in an identifier and advances our position in the input string
// 读取从当前位置到下一个非字母字符的字符串，最终指向下一个非字母字符，该非字母字符不会被处理
// 不支持包含数字的变量名
func (l *Lexer) readIdentifier() string {
	position := l.position
	// Keep reading characters until we encounter a non-letter character
	for isLetter(l.ch) {
		l.readChar()
	}
	// Return the substring of the input string from position to readPosition
	return l.input[position:l.position]
}

// 不支持小数等
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// NextToken returns the next token in the input string
// 处理**当前**的字符，尝试转换成token，转换之后，将会读取下一个字符
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	// Skip over any whitespace characters
	l.skipWhitespace()
	switch l.ch {
	// Operators
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
		// 两个字符的token
		// ==
		if l.peekChar() == '=' {
			l.readChar()
			tok.Literal = "=="
			tok.Type = token.EQ
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		tok = newToken(token.BANG, l.ch)
		// 两个字符的token
		// !=
		if l.peekChar() == '=' {
			l.readChar()
			tok.Literal = "!="
			tok.Type = token.NOT_EQ
		}
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	// Delimiters
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else {
			if isDigit(l.ch) {
				tok.Type = token.INT
				tok.Literal = l.readNumber()
				return tok
			} else {
				tok = newToken(token.ILLEGAL, l.ch)
			}
		}

	}
	// 移动到下一个字符
	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	// We only support ASCII characters for now
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

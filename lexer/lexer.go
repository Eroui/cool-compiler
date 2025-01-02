package lexer

import (
	"bufio"
	"fmt"
	"io"
	"unicode"
)

type TokenType int

// The list of token types
const (
	EOF TokenType = iota
	ERROR

	// Keywords
	CLASS
	IF
	ELSE
	FI
	THEN
	WHILE
	CASE
	ESAC

	// Identifiers
	TYPEID
	OBJECTID

	// Operators
	ASSIGN // <-
	LT     // <
	LE     // <=
	EQ     // =
	PLUS   // +
	MINUS  // -
	TIMES  // *
	DIVIDE // /
	LPAREN // (
	RPAREN // )
	LBRACE // {
	RBRACE // }
	SEMI   // ;
)

// Token represents a lexical token with its type, value, and position.
type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Column  int
}

// Lexer is the lexical analyzer.
type Lexer struct {
	reader *bufio.Reader
	line   int
	column int
	char   rune
}

// NewLexer creates a new lexer from an io.Reader
func NewLexer(reader io.Reader) *Lexer {
	l := &Lexer{
		reader: bufio.NewReader(reader),
		line:   1,
		column: 0,
		char:   ' ',
	}
	return l
}

// readChar reads the next character from the input.
func (l *Lexer) readChar() {
	var err error
	l.char, _, err = l.reader.ReadRune()
	if err != nil {
		l.char = 0 // EOF
	}

	l.column++
	if l.char == '\n' {
		l.line++
		l.column = 0
	}
}

// peekChar returns the next character without advancing the stream.
func (l *Lexer) peekChar() rune {
	char, _, err := l.reader.ReadRune()
	if err != nil {
		return 0
	}
	l.reader.UnreadRune()
	return char
}

// skipWhiteSpace whitespace characters.
func (l *Lexer) skipWhiteSpace() {
	for unicode.IsSpace(l.char) {
		l.readChar()
	}
}

func (l *Lexer) NextToken() Token {
	l.skipWhiteSpace()

	tok := Token{
		Line:   l.line,
		Column: l.column,
	}

	switch {
	case l.char == 0:
		tok.Type = EOF
		tok.Literal = ""
	case l.char == '(':
		tok.Type = LPAREN
		tok.Literal = "("
		l.readChar()
	case l.char == ')':
		tok.Type = RPAREN
		tok.Literal = ")"
		l.readChar()
	case l.char == '{':
		tok.Type = LBRACE
		tok.Literal = "{"
		l.readChar()
	case l.char == '}':
		tok.Type = RBRACE
		tok.Literal = "}"
		l.readChar()
	case l.char == ';':
		tok.Type = SEMI
		tok.Literal = ";"
		l.readChar()
	case l.char == '+':
		tok.Type = PLUS
		tok.Literal = "+"
		l.readChar()
	case l.char == '*':
		tok.Type = TIMES
		tok.Literal = "*"
		l.readChar()
	case l.char == '-':
		tok.Type = MINUS
		tok.Literal = "-"
		l.readChar()
	// This could be a comment or a divide
	// TODO: add support for Multi line comment
	case l.char == '/':
		if l.peekChar() == '/' {
			// This is a single line comment
			for l.char != '\n' && l.char != 0 {
				l.readChar()
			}
			return l.NextToken() // Skip the comment and get the next token
		} else {
			tok.Type = DIVIDE
			tok.Literal = "/"
			l.readChar()
		}
	case l.char == '=':
		tok.Type = EQ
		tok.Literal = "="
		l.readChar()
	// Could be LT, LE, or ASSIGN
	case l.char == '<':
		if l.peekChar() == '-' {
			tok.Type = ASSIGN
			tok.Literal = "<-"
			l.readChar()
			l.readChar()
		} else if l.peekChar() == '=' {
			tok.Type = LE
			tok.Literal = "<="
			l.readChar()
			l.readChar()
		} else {
			tok.Type = LT
			tok.Literal = "<"
			l.readChar()
			l.readChar()
		}

	default:
		tok.Type = ERROR
		tok.Literal = fmt.Sprintf("Unexpected character: %c", l.char)
		l.readChar()
	}

	return tok
}

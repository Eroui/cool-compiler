package lexer

type TokenType int

// The list of token types
const (
	EOF TokenType = iota

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

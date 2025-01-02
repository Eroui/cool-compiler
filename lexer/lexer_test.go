package lexer

import (
	"strings"
	"testing"
)

func TestNextToken(t *testing.T) {
	tests := []struct {
		input             string
		expectedTokenType []TokenType
		expectedLiteral   []string
	}{
		{
			"class Main {};",
			[]TokenType{CLASS, TYPEID, LBRACE, RBRACE, SEMI, EOF},
			[]string{"class", "Main", "{", "}", ";", ""},
		},
		{
			"x <- true;// One line comment\nx <- false;",
			[]TokenType{OBJECTID, ASSIGN, BOOL_CONST, SEMI, OBJECTID, ASSIGN, BOOL_CONST, SEMI, EOF},
			[]string{"x", "<-", "true", ";", "x", "<-", "false", ";", ""},
		},
		{
			"_a <- 0; b   <- _a <= \"1\\n\";",
			[]TokenType{OBJECTID, ASSIGN, INT_CONST, SEMI, OBJECTID, ASSIGN, OBJECTID, LE, STR_CONST, SEMI, EOF},
			[]string{"_a", "<-", "0", ";", "b", "<-", "_a", "<=", "1\n", ";", ""},
		},
	}

	for _, tt := range tests {
		l := NewLexer(strings.NewReader(tt.input))
		for i, expTType := range tt.expectedTokenType {
			tok := l.NextToken()

			if tok.Type != expTType {
				t.Fatalf("[%q]: Wrong token type %d-th Token. expected=%d, got %d", tt.input, i, expTType, tok.Type)
			}

			if tok.Literal != tt.expectedLiteral[i] {
				t.Fatalf("[%q]: Wrong literal at test %d-it Token. expected=%q, got %q", tt.input, i, tt.expectedLiteral[i], tok.Literal)
			}
		}
	}
}

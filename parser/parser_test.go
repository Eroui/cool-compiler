package parser

import (
	"cool-compiler/lexer"
	"strings"
	"testing"
)

func newParserFromInput(input string) *Parser {
	l := lexer.NewLexer(strings.NewReader(input))
	return New(l)
}

func TestClassParser(t *testing.T) {
}

func TestFormalParsing(t *testing.T) {
	tests := []struct {
		input         string
		expectedNames []string
		expectedTypes []string
	}{
		{
			input:         "var1:Integer",
			expectedNames: []string{"var1"},
			expectedTypes: []string{"Integer"},
		},
		{
			input:         "var1:Integer,var2:Boolean,var3:String",
			expectedNames: []string{"var1", "var2", "var3"},
			expectedTypes: []string{"Integer", "Boolean", "String"},
		},
	}

	for _, tt := range tests {
		parser := newParserFromInput(tt.input)
		formals := parser.parseFormals()

		if len(formals) != len(tt.expectedNames) {
			t.Fatalf("[%q]: expected %d formals got %d: %v", tt.input, len(tt.expectedNames), len(formals), formals)
		}

		for i, formal := range formals {
			if formal.Name.Value != tt.expectedNames[i] {
				t.Fatalf("[%q]: expected formal name to be %q got %q", tt.input, tt.expectedNames[i], formal.Name.Value)
			}
			if formal.TypeDecl.Value != tt.expectedTypes[i] {
				t.Fatalf("[%q]: expected formal type to be %q got %q", tt.input, tt.expectedNames[i], formal.Name.Value)
			}
		}

	}
}

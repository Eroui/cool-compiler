package parser

import (
	"cool-compiler/ast"
	"cool-compiler/lexer"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  lexer.Token
	peekToken lexer.Token
	errors    []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program { return nil }

func (p *Parser) parseClass() *ast.Class { return nil }

func (p *Parser) parseFeature() *ast.Feature { return nil }

func (p *Parser) parseMethod() *ast.Method { return nil }

func (p *Parser) parseAttribute() *ast.Attribute { return nil }

func (p *Parser) parseExpression() *ast.Expression { return nil }

package parser

import (
	"cool-compiler/ast"
	"cool-compiler/lexer"
	"fmt"
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

func (p *Parser) curTokenIs(t lexer.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t lexer.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t lexer.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	p.peekError(t)
	return false
}

func (p *Parser) peekError(t lexer.TokenType) {
	p.errors = append(p.errors, fmt.Sprintf("Expected next token to be %v, got %v", t, p.peekToken.Type))
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Classes = []ast.Class{}

	for p.curToken.Type != lexer.EOF {
		if p.curToken.Type != lexer.CLASS {
			class := p.parseClass()
			program.Classes = append(program.Classes, *class)
		} else {
			p.errors = append(p.errors, fmt.Sprintf("Unexpected token: %v, Expected 'class'", p.curToken.Literal))
			return nil
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseClass() *ast.Class {
	class := &ast.Class{Token: p.curToken}

	if !p.expectPeek(lexer.TYPEID) {
		return nil
	}

	class.Name = &ast.TypeIdentifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}

	if !p.expectPeek(lexer.LBRACE) {
		return nil
	}

	for !p.peekTokenIs(lexer.RBRACE) && !p.peekTokenIs(lexer.EOF) {
		p.nextToken()
		feature := p.parseFeature()
		if feature != nil {
			class.Features = append(class.Features, feature)
		}
	}

	if !p.expectPeek(lexer.RBRACE) {
		return nil
	}

	if !p.expectPeek(lexer.SEMI) {
		return nil
	}

	return class
}

func (p *Parser) parseFeature() *ast.Feature { return nil }

func (p *Parser) parseMethod() *ast.Method { return nil }

func (p *Parser) parseAttribute() *ast.Attribute { return nil }

func (p *Parser) parseExpression() *ast.Expression { return nil }

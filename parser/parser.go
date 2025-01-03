package parser

import (
	"cool-compiler/ast"
	"cool-compiler/lexer"
	"fmt"
	"strconv"
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

func (p *Parser) expectAndPeek(t lexer.TokenType) bool {
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

	if !p.expectAndPeek(lexer.TYPEID) {
		return nil
	}

	class.Name = &ast.TypeIdentifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}

	if !p.expectAndPeek(lexer.LBRACE) {
		return nil
	}

	for !p.peekTokenIs(lexer.RBRACE) && !p.peekTokenIs(lexer.EOF) {
		p.nextToken()
		feature := p.parseFeature()
		if feature != nil {
			class.Features = append(class.Features, feature)
		}
	}

	if !p.expectAndPeek(lexer.RBRACE) {
		return nil
	}

	if !p.expectAndPeek(lexer.SEMI) {
		return nil
	}

	return class
}

func (p *Parser) parseFeature() ast.Feature {
	if p.peekTokenIs(lexer.LPAREN) {
		return p.parseMethod()
	}
	return p.parseAttribute()
}

func (p *Parser) parseMethod() *ast.Method {
	method := &ast.Method{Token: p.curToken}
	method.Name = &ast.ObjectIdentifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}

	if !p.expectAndPeek(lexer.LPAREN) {
		return nil
	}

	method.Formals = []*ast.Formal{}
	if !p.peekTokenIs(lexer.RPAREN) {
		p.nextToken()
		method.Formals = p.parseFormals()
	}

	if !p.expectAndPeek(lexer.RPAREN) {
		return nil
	}

	if !p.expectAndPeek(lexer.SEMI) {
		return nil
	}

	if !p.expectAndPeek(lexer.TYPEID) {
		return nil
	}

	method.TypeDecl = &ast.TypeIdentifier{
		Token: p.curToken, Value: p.curToken.Literal,
	}

	if !p.expectAndPeek(lexer.LBRACE) {
		return nil
	}

	p.nextToken()
	method.Expression = p.parseExpression()

	if !p.expectAndPeek(lexer.RBRACE) {
		return nil
	}

	if !p.expectAndPeek(lexer.SEMI) {
		return nil
	}

	return method
}

func (p *Parser) parseAttribute() *ast.Attribute {
	attr := &ast.Attribute{Token: p.curToken}
	attr.Name = &ast.ObjectIdentifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectAndPeek(lexer.COLON) {
		return nil
	}

	if !p.expectAndPeek(lexer.TYPEID) {
		return nil
	}

	attr.TypeDecl = &ast.TypeIdentifier{Token: p.curToken, Value: p.curToken.Literal}

	if p.peekTokenIs(lexer.ASSIGN) {
		p.nextToken()
		p.nextToken()
		attr.Expression = p.parseExpression()
	}

	if !p.expectAndPeek(lexer.SEMI) {
		return nil
	}

	return attr
}

func (p *Parser) parseFormals() []*ast.Formal {
	formals := []*ast.Formal{}

	for {
		if !p.curTokenIs(lexer.OBJECTID) {
			p.errors = append(p.errors, fmt.Sprintf("expected formal parameter name, got %v", p.curToken.Literal))
			return nil
		}

		formal := &ast.Formal{
			Token: p.curToken,
			Name:  &ast.ObjectIdentifier{Token: p.curToken, Value: p.curToken.Literal},
		}

		if !p.expectAndPeek(lexer.COLON) {
			return nil
		}

		if !p.expectAndPeek(lexer.TYPEID) {
			return nil
		}

		formal.TypeDecl = &ast.TypeIdentifier{Token: p.curToken, Value: p.curToken.Literal}
		formals = append(formals, formal)

		if p.peekTokenIs(lexer.COMMA) {
			p.nextToken()
			p.nextToken()
		} else {
			break
		}
	}

	return formals
}

func (p *Parser) parseExpression() ast.Expression {
	switch p.curToken.Type {
	case lexer.INT_CONST:
		return p.parseIntegerLiteral()
	case lexer.STR_CONST:
		return p.parseStringLiteral()
	case lexer.BOOL_CONST:
		return p.parseBooleanLiteral()
	case lexer.OBJECTID:
		if p.peekTokenIs(lexer.ASSIGN) {
			return p.parseAssignment()
		}
		return p.parseIdentifier()
	case lexer.IF:
		return p.parseIfExpression()
	}

	return nil
}

func (p *Parser) parseIntegerLiteral() *ast.IntegerLiteral {
	lit := &ast.IntegerLiteral{Token: p.curToken}

	value, err := strconv.ParseInt(p.curToken.Literal, 10, 64)
	if err != nil {
		p.errors = append(p.errors, fmt.Sprintf("could not parse %q as integer", p.curToken.Literal))
		return nil
	}

	lit.Value = value
	p.nextToken()
	return lit
}

func (p *Parser) parseStringLiteral() *ast.StringLiteral {
	lit := &ast.StringLiteral{Token: p.curToken, Value: p.curToken.Literal}
	p.nextToken()
	return lit
}

func (p *Parser) parseBooleanLiteral() *ast.BooleanLiteral {
	lit := &ast.BooleanLiteral{Token: p.curToken, Value: p.curToken.Literal == "true"}
	p.nextToken()
	return lit
}

func (p *Parser) parseAssignment() *ast.Assignment {
	assign := &ast.Assignment{Token: p.curToken}
	assign.Identifier = &ast.ObjectIdentifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectAndPeek(lexer.ASSIGN) {
		return nil
	}

	p.nextToken()

	assign.Expression = p.parseExpression()

	if p.peekTokenIs(lexer.SEMI) {
		p.nextToken()
	}

	return assign
}

func (p *Parser) parseIdentifier() *ast.ObjectIdentifier {
	identifier := &ast.ObjectIdentifier{Token: p.curToken, Value: p.curToken.Literal}
	p.nextToken()
	return identifier
}

func (p *Parser) parseIfExpression() *ast.IfExpression {
	ie := &ast.IfExpression{Token: p.curToken}
	p.nextToken()
	ie.Condition = p.parseExpression()

	if !p.expectAndPeek(lexer.THEN) {
		return nil
	}

	p.nextToken()
	ie.Consequence = p.parseExpression()

	if !p.expectAndPeek(lexer.THEN) {
		return nil
	}

	p.nextToken()
	ie.Alternative = p.parseExpression()

	if !p.expectAndPeek(lexer.FI) {
		return nil
	}

	if !p.expectAndPeek(lexer.SEMI) {
		return nil
	}

	return ie
}

func (p *Parser) parseWhileExpression() *ast.WhileExpression {
	we := &ast.WhileExpression{Token: p.curToken}

	p.nextToken()
	we.Condition = p.parseExpression()

	if !p.expectAndPeek(lexer.LOOP) {
		return nil
	}

	p.nextToken()

	we.Body = p.parseExpression()

	if !p.expectAndPeek(lexer.POOL) {
		return nil
	}

	if !p.expectAndPeek(lexer.SEMI) {
		return nil
	}

	return we
}

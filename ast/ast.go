package ast

import (
	"cool-compiler/lexer"
)

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Feature interface {
	Node
	featureNode()
}

type TypeIdentifier struct {
	Token lexer.Token
	Value string
}

func (ti *TypeIdentifier) TokenLiteral() string { return ti.Token.Literal }

type ObjectIdentifier struct {
	Token lexer.Token
	Value string
}

func (oi *ObjectIdentifier) TokenLiteral() string { return oi.Token.Literal }
func (oi *ObjectIdentifier) expressionNode()      {}

type Program struct {
	Classes []Class
}

func (p *Program) TokenLiteral() string {
	if len(p.Classes) > 0 {
		return p.Classes[0].TokenLiteral()
	}
	return ""
}

type Class struct {
	Token    lexer.Token
	Name     *TypeIdentifier
	Features []Feature
	Filename string
}

func (c *Class) TokenLiteral() string { return c.Token.Literal }

type Method struct {
	Token      lexer.Token
	Name       *ObjectIdentifier
	Formals    []*Formal
	TypeDecl   *TypeIdentifier
	Expression Expression
}

func (m *Method) TokenLiteral() string { return m.Token.Literal }
func (m *Method) featureNode()         {}

type Formal struct {
	Token    lexer.Token
	Name     *ObjectIdentifier
	TypeDecl *TypeIdentifier
}

func (f *Formal) TokenLiteral() string { return f.Token.Literal }

type Attribute struct {
	Token      lexer.Token
	Name       *ObjectIdentifier
	TypeDecl   *TypeIdentifier
	Expression Expression
}

func (a *Attribute) TokenLiteral() string { return a.Token.Literal }
func (a *Attribute) featureNode()         {}

type IntegerLiteral struct {
	Token lexer.Token
	Value int64
}

func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) expressionNode()      {}

type StringLiteral struct {
	Token lexer.Token
	Value string
}

func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }
func (sl *StringLiteral) expressionNode()      {}

type BooleanLiteral struct {
	Token lexer.Token
	Value bool
}

func (bl *BooleanLiteral) TokenLiteral() string { return bl.Token.Literal }
func (bl *BooleanLiteral) expressionNode()      {}

type Assignment struct {
	Token      lexer.Token
	Identifier *ObjectIdentifier
	Expression Expression
}

func (a *Assignment) TokenLiteral() string { return a.Token.Literal }
func (a *Assignment) expressionNode()      {}

type BinaryExpression struct {
	Token    lexer.Token
	Left     Expression
	Right    Expression
	Operator string
}

func (be *BinaryExpression) TokenLiteral() string { return be.Token.Literal }
func (bl *BinaryExpression) expressionNode()      {}

type UnaryExpression struct {
	Token    lexer.Token
	Right    Expression
	Operator string
}

func (uexp *UnaryExpression) TokenLiteral() string { return uexp.Token.Literal }
func (uexp *UnaryExpression) expressionNode()      {}

type IfExpression struct {
	Token       lexer.Token
	Condition   Expression
	Consequence Expression
	Alternative Expression
}

func (ife *IfExpression) TokenLiteral() string { return ife.Token.Literal }
func (ife *IfExpression) expressionNode()      {}

type WhileExpression struct {
	Token     lexer.Token
	Condition Expression
	Body      Expression
}

func (we *WhileExpression) TokenLiteral() string { return we.Token.Literal }
func (we *WhileExpression) expressionNode()      {}

type BlockExpression struct {
	Token      lexer.Token
	Expression []Expression
}

func (be *BlockExpression) TokenLiteral() string { return be.Token.Literal }
func (be *BlockExpression) expressionNode()      {}

type NewExpression struct {
	Token lexer.Token
	Type  *TypeIdentifier
}

func (ne *NewExpression) TokenLiteral() string { return ne.Token.Literal }
func (ne *NewExpression) expressionNode()      {}

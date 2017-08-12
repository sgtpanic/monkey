package ast

import "monkey/token"

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

// Program is a node because it implements Node
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// LetStatement is a Statement because it implements
// Node and Statement (TokenLiteral and statementNode)
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// Identifier is an Expression because it implements
// TokenLiteral and expressionNode
type Identifier struct {
	Token token.Token
	Value string
}

func (l *Identifier) expressionNode() {}
func (l *Identifier) TokenLiteral() string {
	return l.Token.Literal
}

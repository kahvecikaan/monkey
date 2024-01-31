package ast

import "monkey/token"

// difference between expression and statement: an expression produces a value, a statement does not
// e.g. 5 + 5 is an expression, let x = 5 is a statement

type Node interface {
	TokenLiteral() string // returns the literal value of the token (used only for debugging and testing)
}
type Statement interface {
	Node
	statementNode() // dummy method to distinguish statements from expressions
}

type Expression interface {
	Node
	expressionNode() // dummy method to distinguish statements from expressions
}

type Program struct {
	Statements []Statement // a program is a sequence of statements
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// LetStatement represents a let statement. It consists of a token (the LET token), a name (the identifier that comes
// after the LET token), and an expression that the variable should be bound to (the expression that comes after the
// identifier).
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier // the name of the variable
	Value Expression  // the expression that the variable should be bound to
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string      // the value of the identifier
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

type ReturnStatement struct {
	Token       token.Token // the token.RETURN token
	ReturnValue Expression  // the expression that the return value should be bound to
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

package ast

import (
	"bytes"
	"monkey/token"
)

// difference between expression and statement: an expression produces a value, a statement does not
// e.g. 5 + 5 is an expression, let x = 5 is a statement

type Node interface {
	TokenLiteral() string // returns the literal value of the token (used only for debugging and testing)
	String() string       // returns a string representation of the node (used only for debugging and testing)
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

func (p *Program) String() string { // print the AST
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
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
func (ls *LetStatement) String() string { // print the AST
	var out bytes.Buffer
	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")
	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string      // the value of the identifier
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

type ReturnStatement struct {
	Token       token.Token // the token.RETURN token
	ReturnValue Expression  // the expression that the return value should be bound to
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) String() string { // print the AST
	var out bytes.Buffer
	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

// We need expression statements because, in Monkey the following is legal:
// let x = 5;
// x + 10;
// In other words, we can have expressions that don't produce any value. We call them expression statements.
// In the above example, the expression x + 10 doesn't produce any value. It's not bound to any variable, and it's not
// the return value of a function. It's just an expression that doesn't produce any value.

type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression  //	 the expression itself
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string { // print the AST
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

type IntegerLiteral struct {
	Token token.Token
	Value int64 // note that it's not an string, but int64
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }

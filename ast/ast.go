package ast

import (
	"sputnik/token"

	"bytes"
)

//Everything in AST is a node.
//There are two types of nodes; statements and expressions.
type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

type Identifier struct {
	Token token.Token
	Value string
}

//This has both an identifier and an expression. The expression is for when the identifier is "used"
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

//A return statement consists of two elements, the return token and the expression that will be returned.
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

//The reason behind creating an "Expression statement" is so that we can evaluate expressions without needing a let or return keyword.
//This is usually constrained to just scripting languages.
type ExpressionStatement struct {
	Token token.Token
	Expression Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetStatement) String() string {
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

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(",")

	return out.String()
}

func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}

func (i *Identifier) statementNode()       {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string {
	return i.Value
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	 var out bytes.Buffer

	 for _, s := range p.Statements {
	 	 out.WriteString(s.String())
	 }

	 return out.String()
}
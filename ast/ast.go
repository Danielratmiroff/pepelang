package ast

import (
	"bytes"
	"ratmy/token"
	"strings"
)

type Program struct {
	Statements []Statement
}

type Node interface {
	TokenLiteral() string
	String() string
}

// Node Types
type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// Statement Types
type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

type BlockStatement struct {
	Token      token.Token // the { token
	Statements []Statement
}

type LetStatement struct {
	Token token.Token // the Let token
	Name  *Identifier
	Value Expression
}

type ArrayLiteral struct {
	Token    token.Token // the [ token
	Elements []Expression
}

type StringLiteral struct {
	Token token.Token // the Return token
	Value string
}

type ReturnStatement struct {
	Token       token.Token // the Return token
	ReturnValue Expression
}

type FunctionLiteral struct {
	Token      token.Token // the Function token
	Parameters []*Identifier
	Body       *BlockStatement
}

type CallExpression struct {
	Token     token.Token // the "("  token
	Function  Expression  // Identifier or Function Literal
	Arguments []Expression
}

type IfExpression struct {
	Token       token.Token // the If token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

type Boolean struct {
	Token token.Token // the Bool token
	Value bool
}

type Identifier struct {
	Token token.Token // the IDENT token
	Value string
}

type IntegerLiteral struct {
	Token token.Token // the IDENT token
	Value int64
}

type PrefixExpression struct {
	Token    token.Token // the prefix tokens: ! or -
	Operator string
	Right    Expression
}

type IndexExpression struct {
	Token token.Token // the [ token
	Left  Expression
	Index Expression
}

type InfixExpression struct {
	Token    token.Token // the infix tokens: *, +, ==, etc
	Left     Expression
	Operator string
	Right    Expression
}

// AST Definitions
func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

func (al *ArrayLiteral) expressionNode()      {}
func (al *ArrayLiteral) TokenLiteral() string { return al.Token.Literal }
func (al *ArrayLiteral) String() string {
	var out bytes.Buffer

	elements := []string{}
	for _, el := range al.Elements {
		elements = append(elements, el.String())
	}
	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

func (ls *LetStatement) statementNode() {}
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
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")
	return out.String()

}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

func (ie *IfExpression) expressionNode()      {}
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString("else")
		out.WriteString(ie.Alternative.String())
	}
	return out.String()
}

func (ie *IndexExpression) expressionNode()      {}
func (ie *IndexExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IndexExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("])")

	return out.String()
}

func (fl *FunctionLiteral) expressionNode()      {}
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer
	params := []string{}

	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())

	return out.String()
}

func (ce *CallExpression) expressionNode()      {}
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CallExpression) String() string {
	var out bytes.Buffer
	args := []string{}

	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}

func (b *Boolean) expressionNode()      {}
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) String() string       { return b.Token.Literal }

func (i *Identifier) expressionNode()      {}
func (i *Identifier) String() string       { return i.Value }
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

func (sl *StringLiteral) expressionNode()      {}
func (sl *StringLiteral) String() string       { return sl.Token.Literal }
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) String() string       { return il.TokenLiteral() }
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }

func (pe *PrefixExpression) expressionNode() {}
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }

func (oe *InfixExpression) expressionNode() {}
func (oe *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(oe.Left.String())
	out.WriteString(" " + oe.Operator + " ")
	out.WriteString(oe.Right.String())
	out.WriteString(")")

	return out.String()
}
func (oe *InfixExpression) TokenLiteral() string { return oe.Token.Literal }

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer
	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

// Node's Methods
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

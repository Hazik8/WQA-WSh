package parser

import "windroid/wqa/internal/lexer"

type Node interface {
	node()
}

type Statement interface {
	Node
	statement()
}

type Expression interface {
	Node
	expression()
}

type Program struct {
	Statements []Statement
}

func (Program) node() {}

type SetStatement struct {
	Name  lexer.Token
	Value Expression
}

func (SetStatement) node()      {}
func (SetStatement) statement() {}

type PrintStatement struct {
	Value Expression
}

func (PrintStatement) node()      {}
func (PrintStatement) statement() {}

type BinaryStatement struct {
	Operator lexer.TokenType
	Left     Expression
	Right    Expression
	Result   lexer.Token
}

func (BinaryStatement) node()      {}
func (BinaryStatement) statement() {}

type IdentifierExpression struct {
	Token lexer.Token
}

func (IdentifierExpression) node()       {}
func (IdentifierExpression) expression() {}

type NumberExpression struct {
	Token lexer.Token
}

func (NumberExpression) node()       {}
func (NumberExpression) expression() {}

type StringExpression struct {
	Token lexer.Token
}

func (StringExpression) node()       {}
func (StringExpression) expression() {}

type IfStatement struct {
	Condition Expression
	Then      []Statement
	Else      []Statement
}

func (IfStatement) node()      {}
func (IfStatement) statement() {}

type ComparisonExpression struct {
	Left     Expression
	Operator lexer.TokenType
	Right    Expression
}

func (ComparisonExpression) node()       {}
func (ComparisonExpression) expression() {}

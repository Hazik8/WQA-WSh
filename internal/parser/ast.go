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

type LoopStatement struct {
	Condition Expression
	Body      []Statement
}

func (LoopStatement) node()      {}
func (LoopStatement) statement() {}

type RepeatStatement struct {
	Count Expression
	Body  []Statement
}

func (RepeatStatement) node()      {}
func (RepeatStatement) statement() {}

type FuncStatement struct {
	Name string
	Body []Statement
}

func (FuncStatement) node()      {}
func (FuncStatement) statement() {}

type CallStatement struct {
	Name   string
	Result string
}

func (CallStatement) node()      {}
func (CallStatement) statement() {}

type GiveStatement struct {
	Value Expression
}

func (GiveStatement) node()      {}
func (GiveStatement) statement() {}

type InputStatement struct {
	Name string
}

func (InputStatement) node()      {}
func (InputStatement) statement() {}

type ClearStatement struct{}

func (ClearStatement) node()      {}
func (ClearStatement) statement() {}

type WaitStatement struct {
	Duration Expression
}

func (WaitStatement) node()      {}
func (WaitStatement) statement() {}

type ExitStatement struct{}

func (ExitStatement) node()      {}
func (ExitStatement) statement() {}

type TimeStatement struct{}

func (TimeStatement) node()      {}
func (TimeStatement) statement() {}

type DateStatement struct{}

func (DateStatement) node()      {}
func (DateStatement) statement() {}

type BinaryExpression struct {
	Left     Expression
	Operator lexer.TokenType
	Right    Expression
}

func (BinaryExpression) node()       {}
func (BinaryExpression) expression() {}

type UnaryExpression struct {
	Operator lexer.TokenType
	Right    Expression
}

func (UnaryExpression) node()       {}
func (UnaryExpression) expression() {}

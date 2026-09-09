package compiler

import (
	"fmt"
	"strconv"

	"windroid/wqa/internal/lexer"
	"windroid/wqa/internal/parser"
	"windroid/wqa/internal/runtime"
)

type Compiler struct {
	emitter *Emitter
}

func New() *Compiler {
	return &Compiler{
		emitter: NewEmitter(),
	}
}

func Compile(program *parser.Program) ([]byte, error) {
	c := New()
	return c.Compile(program)
}

func (c *Compiler) Compile(
	program *parser.Program,
) ([]byte, error) {

	for _, statement := range program.Statements {
		if err := c.compileStatement(statement); err != nil {
			return nil, err
		}
	}

	return c.emitter.Bytes(), nil
}

func (c *Compiler) compileStatement(
	statement parser.Statement,
) error {

	switch stmt := statement.(type) {

	case parser.SetStatement:
		return c.compileSet(stmt)

	case parser.PrintStatement:
		return c.compilePrint(stmt)

	case parser.BinaryStatement:
		return c.compileBinary(stmt)

	case parser.IfStatement:
		return c.compileIf(stmt)

	default:
		return fmt.Errorf(
			"unsupported statement: %T",
			statement,
		)
	}
}

func (c *Compiler) compileSet(
	stmt parser.SetStatement,
) error {

	if err := c.compileExpression(stmt.Value); err != nil {
		return err
	}

	c.emitter.Emit(runtime.OP_SET)

	c.emitter.EmitString(
		stmt.Name.Lexeme,
	)

	return nil
}

func (c *Compiler) compilePrint(
	stmt parser.PrintStatement,
) error {

	if err := c.compileExpression(stmt.Value); err != nil {
		return err
	}

	c.emitter.Emit(runtime.OP_PRINT)

	return nil
}

func (c *Compiler) compileBinary(
	stmt parser.BinaryStatement,
) error {

	if err := c.compileExpression(stmt.Left); err != nil {
		return err
	}

	if err := c.compileExpression(stmt.Right); err != nil {
		return err
	}

	switch stmt.Operator {

	case lexer.TokenAdd:
		c.emitter.Emit(runtime.OP_ADD)

	case lexer.TokenSub:
		c.emitter.Emit(runtime.OP_SUB)

	case lexer.TokenMul:
		c.emitter.Emit(runtime.OP_MUL)

	case lexer.TokenDiv:
		c.emitter.Emit(runtime.OP_DIV)

	default:
		return fmt.Errorf(
			"unknown binary operator: %v",
			stmt.Operator,
		)
	}

	c.emitter.Emit(runtime.OP_SET)

	c.emitter.EmitString(
		stmt.Result.Lexeme,
	)

	return nil
}

func (c *Compiler) compileIf(
	stmt parser.IfStatement,
) error {

	condition, ok := stmt.Condition.(parser.ComparisonExpression)

	if !ok {
		return fmt.Errorf(
			"wif requires a comparison expression",
		)
	}

	// Левая часть условия.
	if err := c.compileExpression(condition.Left); err != nil {
		return err
	}

	// Правая часть условия.
	if err := c.compileExpression(condition.Right); err != nil {
		return err
	}

	// OP_IF берёт два значения со стека
	// и получает оператор следующим параметром.
	c.emitter.Emit(runtime.OP_IF)

	switch condition.Operator {

	case lexer.TokenGreater:
		c.emitter.EmitString(">")

	case lexer.TokenLess:
		c.emitter.EmitString("<")

	case lexer.TokenEqualEqual:
		c.emitter.EmitString("==")

	default:
		return fmt.Errorf(
			"unknown comparison operator: %v",
			condition.Operator,
		)
	}

	// THEN
	for _, statement := range stmt.Then {
		if err := c.compileStatement(statement); err != nil {
			return err
		}
	}

	// ELSE
	if len(stmt.Else) > 0 {

		c.emitter.Emit(runtime.OP_ELSE)

		for _, statement := range stmt.Else {
			if err := c.compileStatement(statement); err != nil {
				return err
			}
		}
	}

	// ENDWIF
	c.emitter.Emit(runtime.OP_ENDIF)

	return nil
}

func (c *Compiler) compileExpression(
	expr parser.Expression,
) error {

	switch value := expr.(type) {

	case parser.IdentifierExpression:

		c.emitter.Emit(runtime.OP_LOAD)

		c.emitter.EmitString(
			value.Token.Lexeme,
		)

	case parser.NumberExpression:

		number := value.Token.Lexeme

		if _, err := strconv.Atoi(number); err != nil {
			return fmt.Errorf(
				"invalid number: %s",
				number,
			)
		}

		c.emitter.Emit(runtime.OP_PUSH)

		c.emitter.EmitString(number)

	case parser.StringExpression:

		c.emitter.Emit(runtime.OP_PUSH)

		c.emitter.EmitString(
			value.Token.Lexeme,
		)

	default:

		return fmt.Errorf(
			"unsupported expression: %T",
			expr,
		)
	}

	return nil
}

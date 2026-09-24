package compiler

import (
	"fmt"
	"strconv"

	"windroid/wqa/internal/lexer"
	"windroid/wqa/internal/parser"
	"windroid/wqa/internal/runtime"
)

type Compiler struct {
	emitter       *Emitter
	functions     map[string]int
	functionCalls map[string][]int
}

func New() *Compiler {
	return &Compiler{
		emitter:       NewEmitter(),
		functions:     make(map[string]int),
		functionCalls: make(map[string][]int),
	}
}

func Compile(program *parser.Program) ([]byte, error) {
	c := New()
	return c.Compile(program)
}

func (c *Compiler) Compile(
	program *parser.Program,
) ([]byte, error) {

	// Сначала собираем имена функций.
	for _, statement := range program.Statements {
		if stmt, ok := statement.(parser.FuncStatement); ok {

			if _, exists := c.functions[stmt.Name]; exists {
				return nil, fmt.Errorf(
					"function %q already defined",
					stmt.Name,
				)
			}

			c.functions[stmt.Name] = -1
		}
	}

	// Основная программа.
	for _, statement := range program.Statements {
		if _, ok := statement.(parser.FuncStatement); ok {
			continue
		}

		if err := c.compileStatement(statement); err != nil {
			return nil, err
		}
	}

	// После основной программы выполнение заканчивается.
	c.emitter.Emit(runtime.OP_EXIT)

	// Компилируем функции.
	for _, statement := range program.Statements {
		stmt, ok := statement.(parser.FuncStatement)
		if !ok {
			continue
		}

		address := len(c.emitter.Bytes())
		c.functions[stmt.Name] = address

		for _, bodyStatement := range stmt.Body {
			if err := c.compileStatement(bodyStatement); err != nil {
				return nil, err
			}
		}

		c.emitter.Emit(runtime.OP_RET)
	}

	// Заполняем адреса всех call.
	for name, positions := range c.functionCalls {

		address, exists := c.functions[name]
		if !exists || address < 0 {
			return nil, fmt.Errorf(
				"undefined function: %s",
				name,
			)
		}

		for _, position := range positions {
			c.patchUint32(
				position,
				uint32(address),
			)
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

	case parser.LoopStatement:
		return c.compileLoop(stmt)

	case parser.RepeatStatement:
		return c.compileRepeat(stmt)

	case parser.CallStatement:
		return c.compileCall(stmt)

	case parser.GiveStatement:
		return c.compileGive(stmt)

	case parser.InputStatement:
		return c.compileInput(stmt)

	case parser.ClearStatement:
		return c.compileClear(stmt)

	case parser.WaitStatement:
		return c.compileWait(stmt)

	case parser.ExitStatement:
		return c.compileExit(stmt)

	case parser.TimeStatement:
		return c.compileTime(stmt)

	case parser.DateStatement:
		return c.compileDate(stmt)

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

func (c *Compiler) compileTime(_ parser.TimeStatement) error {
	c.emitter.Emit(runtime.OP_TIME)
	return nil
}

func (c *Compiler) compileDate(_ parser.DateStatement) error {
	c.emitter.Emit(runtime.OP_DATE)
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
		if _, err := strconv.ParseFloat(value.Token.Lexeme, 64); err != nil {
			return fmt.Errorf("invalid number: %s", value.Token.Lexeme)
		}

		c.emitter.Emit(runtime.OP_PUSH)
		c.emitter.EmitString(value.Token.Lexeme)

	case parser.StringExpression:

		c.emitter.Emit(runtime.OP_PUSH)

		c.emitter.EmitString(
			value.Token.Lexeme,
		)

	case parser.BinaryExpression:
		if err := c.compileExpression(value.Left); err != nil {
			return err
		}

		if err := c.compileExpression(value.Right); err != nil {
			return err
		}

		switch value.Operator {
		case lexer.TokenPlus:
			c.emitter.Emit(runtime.OP_ADD)

		case lexer.TokenMinus:
			c.emitter.Emit(runtime.OP_SUB)

		case lexer.TokenStar:
			c.emitter.Emit(runtime.OP_MUL)

		case lexer.TokenSlash:
			c.emitter.Emit(runtime.OP_DIV)

		case lexer.TokenDoubleSlash:
			c.emitter.Emit(runtime.OP_IDIV)

		case lexer.TokenPercent:
			c.emitter.Emit(runtime.OP_MOD)

		default:
			return fmt.Errorf(
				"unknown expression operator: %v",
				value.Operator,
			)
		}

	case parser.UnaryExpression:
		if err := c.compileExpression(value.Right); err != nil {
			return err
		}

		switch value.Operator {
		case lexer.TokenMinus:
			c.emitter.Emit(runtime.OP_NEG)

		default:
			return fmt.Errorf(
				"unknown unary operator: %v",
				value.Operator,
			)
		}

	default:

		return fmt.Errorf(
			"unsupported expression: %T",
			expr,
		)
	}

	return nil
}

func (c *Compiler) compileLoop(
	stmt parser.LoopStatement,
) error {

	condition, ok := stmt.Condition.(parser.ComparisonExpression)

	if !ok {
		return fmt.Errorf(
			"wloop requires a comparison expression",
		)
	}

	// Запоминаем начало условия.
	loopStart := len(c.emitter.Bytes())

	// Левая часть.
	if err := c.compileExpression(condition.Left); err != nil {
		return err
	}

	// Правая часть.
	if err := c.compileExpression(condition.Right); err != nil {
		return err
	}

	// Проверка условия.
	c.emitter.Emit(runtime.OP_LOOP)

	switch condition.Operator {

	case lexer.TokenGreater:
		c.emitter.EmitString(">")

	case lexer.TokenLess:
		c.emitter.EmitString("<")

	case lexer.TokenEqualEqual:
		c.emitter.EmitString("==")

	default:
		return fmt.Errorf(
			"unknown loop comparison operator: %v",
			condition.Operator,
		)
	}

	// Место для адреса выхода из цикла.
	exitPosition := len(c.emitter.Bytes())

	c.emitter.EmitBytes([]byte{0, 0, 0, 0})

	// Тело цикла.
	for _, statement := range stmt.Body {
		if err := c.compileStatement(statement); err != nil {
			return err
		}
	}

	// Прыжок обратно к условию.
	c.emitter.Emit(runtime.OP_JUMP)

	// Адрес начала условия.
	c.emitter.EmitBytes(uint32Bytes(uint32(loopStart)))

	// Адрес сюда будет подставлен позже.
	loopEnd := len(c.emitter.Bytes())

	// Конец цикла.
	c.patchUint32(
		exitPosition,
		uint32(loopEnd),
	)

	c.emitter.Emit(runtime.OP_ENDLOOP)

	return nil
}

func uint32Bytes(value uint32) []byte {
	return []byte{
		byte(value >> 24),
		byte(value >> 16),
		byte(value >> 8),
		byte(value),
	}
}

func (c *Compiler) patchUint32(
	position int,
	value uint32,
) {
	code := c.emitter.Bytes()

	code[position] = byte(value >> 24)
	code[position+1] = byte(value >> 16)
	code[position+2] = byte(value >> 8)
	code[position+3] = byte(value)
}

func (c *Compiler) compileRepeat(
	stmt parser.RepeatStatement,
) error {

	if err := c.compileExpression(stmt.Count); err != nil {
		return err
	}

	c.emitter.Emit(runtime.OP_REPEAT)

	exitPosition := len(c.emitter.Bytes())

	// Место под адрес выхода.
	c.emitter.EmitBytes([]byte{0, 0, 0, 0})

	// Тело цикла.
	for _, statement := range stmt.Body {
		if err := c.compileStatement(statement); err != nil {
			return err
		}
	}

	// Следующая итерация или выход.
	c.emitter.Emit(runtime.OP_ENDREPEAT)

	// Адрес после цикла.
	exitAddress := uint32(len(c.emitter.Bytes()))

	c.patchUint32(
		exitPosition,
		exitAddress,
	)

	return nil
}

func (c *Compiler) compileFunc(
	stmt parser.FuncStatement,
) error {

	for _, statement := range stmt.Body {
		if err := c.compileStatement(statement); err != nil {
			return err
		}
	}

	return nil
}

func (c *Compiler) compileCall(
	stmt parser.CallStatement,
) error {

	c.emitter.Emit(runtime.OP_CALL)

	position := len(c.emitter.Bytes())

	c.emitter.EmitBytes([]byte{0, 0, 0, 0})

	c.functionCalls[stmt.Name] = append(
		c.functionCalls[stmt.Name],
		position,
	)

	// Если указана переменная результата,
	// после возврата сохраняем значение из стека.
	if stmt.Result != "" {
		c.emitter.Emit(runtime.OP_SET)
		c.emitter.EmitString(stmt.Result)
	}

	return nil
}

func (c *Compiler) compileGive(
	stmt parser.GiveStatement,
) error {

	if err := c.compileExpression(stmt.Value); err != nil {
		return err
	}

	c.emitter.Emit(runtime.OP_RET)

	return nil
}

func (c *Compiler) compileInput(
	stmt parser.InputStatement,
) error {

	c.emitter.Emit(runtime.OP_INPUT)
	c.emitter.EmitString(stmt.Name)

	return nil
}

func (c *Compiler) compileClear(
	_ parser.ClearStatement,
) error {

	c.emitter.Emit(runtime.OP_CLEAR)

	return nil
}

func (c *Compiler) compileWait(
	stmt parser.WaitStatement,
) error {

	if err := c.compileExpression(stmt.Duration); err != nil {
		return err
	}

	c.emitter.Emit(runtime.OP_WAIT)

	return nil
}

func (c *Compiler) compileExit(
	_ parser.ExitStatement,
) error {

	c.emitter.Emit(runtime.OP_EXIT)

	return nil
}

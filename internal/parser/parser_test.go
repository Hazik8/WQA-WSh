package parser

import (
	"testing"

	"windroid/wqa/internal/lexer"
)

func parseSource(t *testing.T, source string) *Program {
	t.Helper()

	source = "wqa\n" + source

	tokens, err := lexer.New(source).Tokenize()
	if err != nil {
		t.Fatalf("lexer error: %v", err)
	}

	program, err := New(tokens).Parse()
	if err != nil {
		t.Fatalf("parser error: %v", err)
	}

	return program
}

func TestParseSetNumber(t *testing.T) {
	program := parseSource(
		t,
		"wet x = 10\n",
	)

	if len(program.Statements) != 1 {
		t.Fatalf(
			"expected 1 statement, got %d",
			len(program.Statements),
		)
	}

	statement, ok := program.Statements[0].(SetStatement)
	if !ok {
		t.Fatalf("expected SetStatement")
	}

	if statement.Name.Lexeme != "x" {
		t.Errorf(
			"expected variable x, got %q",
			statement.Name.Lexeme,
		)
	}

	value, ok := statement.Value.(NumberExpression)
	if !ok {
		t.Fatalf("expected NumberExpression")
	}

	if value.Token.Lexeme != "10" {
		t.Errorf(
			"expected number 10, got %q",
			value.Token.Lexeme,
		)
	}
}

func TestParseSetString(t *testing.T) {
	program := parseSource(
		t,
		`wet message = "Hello World"`+"\n",
	)

	statement, ok := program.Statements[0].(SetStatement)
	if !ok {
		t.Fatalf("expected SetStatement")
	}

	value, ok := statement.Value.(StringExpression)
	if !ok {
		t.Fatalf("expected StringExpression")
	}

	if value.Token.Lexeme != "Hello World" {
		t.Errorf(
			"expected Hello World, got %q",
			value.Token.Lexeme,
		)
	}
}

func TestParsePrint(t *testing.T) {
	program := parseSource(
		t,
		`print "Hello"`+"\n",
	)

	statement, ok := program.Statements[0].(PrintStatement)
	if !ok {
		t.Fatalf("expected PrintStatement")
	}

	value, ok := statement.Value.(StringExpression)
	if !ok {
		t.Fatalf("expected StringExpression")
	}

	if value.Token.Lexeme != "Hello" {
		t.Errorf(
			"expected Hello, got %q",
			value.Token.Lexeme,
		)
	}
}

func TestParsePrintIdentifier(t *testing.T) {
	program := parseSource(
		t,
		"print x\n",
	)

	statement, ok := program.Statements[0].(PrintStatement)
	if !ok {
		t.Fatalf("expected PrintStatement")
	}

	value, ok := statement.Value.(IdentifierExpression)
	if !ok {
		t.Fatalf("expected IdentifierExpression")
	}

	if value.Token.Lexeme != "x" {
		t.Errorf(
			"expected x, got %q",
			value.Token.Lexeme,
		)
	}
}

func TestParseAdd(t *testing.T) {
	program := parseSource(
		t,
		"add x y result\n",
	)

	statement, ok := program.Statements[0].(BinaryStatement)
	if !ok {
		t.Fatalf("expected BinaryStatement")
	}

	if statement.Operator != lexer.TokenAdd {
		t.Errorf(
			"expected ADD operator, got %s",
			statement.Operator,
		)
	}

	left, ok := statement.Left.(IdentifierExpression)
	if !ok {
		t.Fatalf("expected identifier as left operand")
	}

	right, ok := statement.Right.(IdentifierExpression)
	if !ok {
		t.Fatalf("expected identifier as right operand")
	}

	if left.Token.Lexeme != "x" {
		t.Errorf(
			"expected left operand x, got %q",
			left.Token.Lexeme,
		)
	}

	if right.Token.Lexeme != "y" {
		t.Errorf(
			"expected right operand y, got %q",
			right.Token.Lexeme,
		)
	}

	if statement.Result.Lexeme != "result" {
		t.Errorf(
			"expected result variable result, got %q",
			statement.Result.Lexeme,
		)
	}
}

func TestParseAllBinaryOperations(t *testing.T) {
	tests := []struct {
		name     string
		source   string
		operator lexer.TokenType
	}{
		{
			name:     "add",
			source:   "add x y result\n",
			operator: lexer.TokenAdd,
		},
		{
			name:     "sub",
			source:   "sub x y result\n",
			operator: lexer.TokenSub,
		},
		{
			name:     "mul",
			source:   "mul x y result\n",
			operator: lexer.TokenMul,
		},
		{
			name:     "div",
			source:   "div x y result\n",
			operator: lexer.TokenDiv,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			program := parseSource(
				t,
				test.source,
			)

			statement, ok := program.Statements[0].(BinaryStatement)
			if !ok {
				t.Fatalf("expected BinaryStatement")
			}

			if statement.Operator != test.operator {
				t.Errorf(
					"expected %s, got %s",
					test.operator,
					statement.Operator,
				)
			}
		})
	}
}

func TestParseMultipleStatements(t *testing.T) {
	source := `wet x = 10
wet y = 20
add x y result
print result
`

	program := parseSource(t, source)

	if len(program.Statements) != 4 {
		t.Fatalf(
			"expected 4 statements, got %d",
			len(program.Statements),
		)
	}
}

func TestParseEmptyProgram(t *testing.T) {
	program := parseSource(
		t,
		"\n\n",
	)

	if len(program.Statements) != 0 {
		t.Fatalf(
			"expected empty program, got %d statements",
			len(program.Statements),
		)
	}
}

func TestParseMissingVariableName(t *testing.T) {
	source := "wet = 10\n"

	tokens, err := lexer.New(source).Tokenize()
	if err != nil {
		t.Fatalf("unexpected lexer error: %v", err)
	}

	_, err = New(tokens).Parse()

	if err == nil {
		t.Fatal("expected parser error")
	}
}

func TestParseMissingEquals(t *testing.T) {
	source := "wet x 10\n"

	tokens, err := lexer.New(source).Tokenize()
	if err != nil {
		t.Fatalf("unexpected lexer error: %v", err)
	}

	_, err = New(tokens).Parse()

	if err == nil {
		t.Fatal("expected parser error")
	}
}

func TestParseMissingExpression(t *testing.T) {
	source := "print\n"

	tokens, err := lexer.New(source).Tokenize()
	if err != nil {
		t.Fatalf("unexpected lexer error: %v", err)
	}

	_, err = New(tokens).Parse()

	if err == nil {
		t.Fatal("expected parser error")
	}
}

func TestParseMissingBinaryOperand(t *testing.T) {
	source := "add x\n"

	tokens, err := lexer.New(source).Tokenize()
	if err != nil {
		t.Fatalf("unexpected lexer error: %v", err)
	}

	_, err = New(tokens).Parse()

	if err == nil {
		t.Fatal("expected parser error")
	}
}

func TestParseUnknownStatement(t *testing.T) {
	source := "hello world\n"

	tokens, err := lexer.New(source).Tokenize()
	if err != nil {
		t.Fatalf("unexpected lexer error: %v", err)
	}

	_, err = New(tokens).Parse()

	if err == nil {
		t.Fatal("expected parser error")
	}
}

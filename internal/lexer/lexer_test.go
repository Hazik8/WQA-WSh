package lexer

import "testing"

func TestLexerBasicCommands(t *testing.T) {
	source := `wet x = 10
print "Hello"
add x y result
`

	lexer := New(source)

	tokens, err := lexer.Tokenize()
	if err != nil {
		t.Fatalf("unexpected lexer error: %v", err)
	}

	expected := []TokenType{
		TokenSet,
		TokenIdentifier,
		TokenEquals,
		TokenNumber,
		TokenNewline,

		TokenPrint,
		TokenString,
		TokenNewline,

		TokenAdd,
		TokenIdentifier,
		TokenIdentifier,
		TokenIdentifier,
		TokenNewline,

		TokenEOF,
	}

	if len(tokens) != len(expected) {
		t.Fatalf(
			"expected %d tokens, got %d",
			len(expected),
			len(tokens),
		)
	}

	for i, expectedType := range expected {
		if tokens[i].Type != expectedType {
			t.Errorf(
				"token %d: expected %s, got %s",
				i,
				expectedType,
				tokens[i].Type,
			)
		}
	}
}

func TestLexerIdentifiers(t *testing.T) {
	source := `myVariable
_test
value123
`

	l := New(source)

	tokens, err := l.Tokenize()
	if err != nil {
		t.Fatalf("unexpected lexer error: %v", err)
	}

	expected := []string{
		"myVariable",
		"_test",
		"value123",
	}

	index := 0

	for _, token := range tokens {
		if token.Type == TokenIdentifier {
			if index >= len(expected) {
				t.Fatalf("unexpected extra identifier: %q", token.Lexeme)
			}

			if token.Lexeme != expected[index] {
				t.Errorf(
					"identifier %d: expected %q, got %q",
					index,
					expected[index],
					token.Lexeme,
				)
			}

			index++
		}
	}

	if index != len(expected) {
		t.Fatalf(
			"expected %d identifiers, got %d",
			len(expected),
			index,
		)
	}
}

func TestLexerNumbers(t *testing.T) {
	source := `123
456
0
`

	l := New(source)

	tokens, err := l.Tokenize()
	if err != nil {
		t.Fatalf("unexpected lexer error: %v", err)
	}

	expected := []string{
		"123",
		"456",
		"0",
	}

	index := 0

	for _, token := range tokens {
		if token.Type != TokenNumber {
			continue
		}

		if index >= len(expected) {
			t.Fatalf("unexpected extra number: %q", token.Lexeme)
		}

		if token.Lexeme != expected[index] {
			t.Errorf(
				"number %d: expected %q, got %q",
				index,
				expected[index],
				token.Lexeme,
			)
		}

		index++
	}

	if index != len(expected) {
		t.Fatalf(
			"expected %d numbers, got %d",
			len(expected),
			index,
		)
	}
}

func TestLexerStrings(t *testing.T) {
	source := `print "Hello World"
print "WQA 1.0"
`

	l := New(source)

	tokens, err := l.Tokenize()
	if err != nil {
		t.Fatalf("unexpected lexer error: %v", err)
	}

	expected := []string{
		"Hello World",
		"WQA 1.0",
	}

	index := 0

	for _, token := range tokens {
		if token.Type != TokenString {
			continue
		}

		if index >= len(expected) {
			t.Fatalf("unexpected extra string: %q", token.Lexeme)
		}

		if token.Lexeme != expected[index] {
			t.Errorf(
				"string %d: expected %q, got %q",
				index,
				expected[index],
				token.Lexeme,
			)
		}

		index++
	}

	if index != len(expected) {
		t.Fatalf(
			"expected %d strings, got %d",
			len(expected),
			index,
		)
	}
}

func TestLexerPositions(t *testing.T) {
	source := `wet x = 10
print "Hello"
`

	l := New(source)

	tokens, err := l.Tokenize()
	if err != nil {
		t.Fatalf("unexpected lexer error: %v", err)
	}

	if tokens[0].Line != 1 {
		t.Errorf(
			"expected first token on line 1, got %d",
			tokens[0].Line,
		)
	}

	if tokens[5].Line != 2 {
		t.Errorf(
			"expected PRINT token on line 2, got %d",
			tokens[5].Line,
		)
	}
}

func TestLexerUnexpectedCharacter(t *testing.T) {
	source := `wet x = @`

	l := New(source)

	_, err := l.Tokenize()

	if err == nil {
		t.Fatal("expected lexer error, got nil")
	}
}

func TestLexerUnterminatedString(t *testing.T) {
	source := `print "Hello`

	l := New(source)

	_, err := l.Tokenize()

	if err == nil {
		t.Fatal("expected unterminated string error, got nil")
	}
}

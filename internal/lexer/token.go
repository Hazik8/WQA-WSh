package lexer

type TokenType string

const (
	TokenEOF TokenType = "EOF"

	TokenIdentifier TokenType = "IDENTIFIER"
	TokenNumber     TokenType = "NUMBER"
	TokenString     TokenType = "STRING"

	TokenWQA TokenType = "WQA"

	TokenSet   TokenType = "SET"
	TokenPrint TokenType = "PRINT"
	TokenAdd   TokenType = "ADD"
	TokenSub   TokenType = "SUB"
	TokenMul   TokenType = "MUL"
	TokenDiv   TokenType = "DIV"

	TokenIf    TokenType = "IF"
	TokenElse  TokenType = "ELSE"
	TokenEndIf TokenType = "ENDIF"

	TokenEquals     TokenType = "="
	TokenEqualEqual TokenType = "=="
	TokenGreater    TokenType = ">"
	TokenLess       TokenType = "<"

	TokenLeftParen  TokenType = "("
	TokenRightParen TokenType = ")"

	TokenNewline TokenType = "NEWLINE"
)

type Token struct {
	Type   TokenType
	Lexeme string
	Line   int
	Column int
}

func NewToken(
	tokenType TokenType,
	lexeme string,
	line int,
	column int,
) Token {
	return Token{
		Type:   tokenType,
		Lexeme: lexeme,
		Line:   line,
		Column: column,
	}
}

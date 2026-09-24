package lexer

import (
	"fmt"
	"strings"
	"unicode"
)

type Lexer struct {
	source   string
	position int
	line     int
	column   int
}

func New(source string) *Lexer {
	return &Lexer{
		source: source,
		line:   1,
		column: 1,
	}
}

func (l *Lexer) Tokenize() ([]Token, error) {
	var tokens []Token

	for !l.isAtEnd() {
		token, err := l.nextToken()
		if err != nil {
			return nil, err
		}

		tokens = append(tokens, token)
	}

	tokens = append(tokens, NewToken(
		TokenEOF,
		"",
		l.line,
		l.column,
	))

	return tokens, nil
}

func (l *Lexer) nextToken() (Token, error) {
	l.skipSpaces()

	if l.isAtEnd() {
		return NewToken(TokenEOF, "", l.line, l.column), nil
	}

	startLine := l.line
	startColumn := l.column
	ch := l.peek()

	if ch == '\n' {
		l.advance()
		return NewToken(TokenNewline, "\n", startLine, startColumn), nil
	}

	if ch == '\r' {
		l.advance()

		if l.peek() == '\n' {
			l.advance()
		}

		return NewToken(TokenNewline, "\n", startLine, startColumn), nil
	}

	if ch == '"' {
		return l.readString(startLine, startColumn)
	}

	switch ch {
	case '=':
		l.advance()

		if l.peek() == '=' {
			l.advance()
			return NewToken(TokenEqualEqual, "==", startLine, startColumn), nil
		}

		return NewToken(TokenEquals, "=", startLine, startColumn), nil

	case '+':
		l.advance()
		return NewToken(TokenPlus, "+", startLine, startColumn), nil

	case '-':
		l.advance()
		return NewToken(TokenMinus, "-", startLine, startColumn), nil

	case '*':
		l.advance()
		return NewToken(TokenStar, "*", startLine, startColumn), nil

	case '/':
		l.advance()

		if l.peek() == '/' {
			l.advance()
			return NewToken(TokenDoubleSlash, "//", startLine, startColumn), nil
		}

		return NewToken(TokenSlash, "/", startLine, startColumn), nil

	case '%':
		l.advance()
		return NewToken(TokenPercent, "%", startLine, startColumn), nil

	case '>':
		l.advance()

		if l.peek() == '=' {
			l.advance()
			return NewToken(TokenGreaterEq, ">=", startLine, startColumn), nil
		}

		return NewToken(TokenGreater, ">", startLine, startColumn), nil

	case '<':
		l.advance()

		if l.peek() == '=' {
			l.advance()
			return NewToken(TokenLessEq, "<=", startLine, startColumn), nil
		}

		return NewToken(TokenLess, "<", startLine, startColumn), nil

	case '!':
		l.advance()

		if l.peek() == '=' {
			l.advance()
			return NewToken(TokenNotEqual, "!=", startLine, startColumn), nil
		}

		return Token{}, fmt.Errorf(
			"lexer error at %d:%d: unexpected character %q",
			startLine,
			startColumn,
			ch,
		)

	case '(':
		l.advance()
		return NewToken(TokenLeftParen, "(", startLine, startColumn), nil

	case ')':
		l.advance()
		return NewToken(TokenRightParen, ")", startLine, startColumn), nil
	}

	if unicode.IsDigit(rune(ch)) {
		return l.readNumber(startLine, startColumn), nil
	}

	if isIdentifierStart(ch) {
		return l.readIdentifier(startLine, startColumn), nil
	}

	return Token{}, fmt.Errorf(
		"lexer error at %d:%d: unexpected character %q",
		startLine,
		startColumn,
		ch,
	)
}

func (l *Lexer) readNumber(line, column int) Token {
	start := l.position
	dotSeen := false

	for !l.isAtEnd() {
		ch := l.peek()

		if ch >= '0' && ch <= '9' {
			l.advance()
			continue
		}

		if ch == '.' && !dotSeen {
			dotSeen = true
			l.advance()
			continue
		}

		break
	}

	return NewToken(
		TokenNumber,
		l.source[start:l.position],
		line,
		column,
	)
}

func (l *Lexer) readIdentifier(line, column int) Token {
	start := l.position

	for !l.isAtEnd() && isIdentifierPart(l.peek()) {
		l.advance()
	}

	lexeme := l.source[start:l.position]

	return NewToken(
		keywordToken(lexeme),
		lexeme,
		line,
		column,
	)
}

func (l *Lexer) readString(line, column int) (Token, error) {
	l.advance()

	var builder strings.Builder

	for !l.isAtEnd() {
		ch := l.peek()

		if ch == '"' {
			l.advance()

			return NewToken(
				TokenString,
				builder.String(),
				line,
				column,
			), nil
		}

		if ch == '\n' || ch == '\r' {
			return Token{}, fmt.Errorf(
				"lexer error at %d:%d: unterminated string",
				line,
				column,
			)
		}

		builder.WriteByte(ch)
		l.advance()
	}

	return Token{}, fmt.Errorf(
		"lexer error at %d:%d: unterminated string",
		line,
		column,
	)
}

func keywordToken(value string) TokenType {
	switch strings.ToLower(value) {
	case "wqa":
		return TokenWQA
	case "wet":
		return TokenSet
	case "print":
		return TokenPrint

	case "add":
		return TokenAdd
	case "sub":
		return TokenSub
	case "mul":
		return TokenMul
	case "div":
		return TokenDiv

	case "wif":
		return TokenIf
	case "else":
		return TokenElse
	case "endwif":
		return TokenEndIf

	case "wloop":
		return TokenLoop
	case "endloop":
		return TokenEndLoop

	case "repeat":
		return TokenRepeat
	case "endrepeat":
		return TokenEndRepeat

	case "wfunc":
		return TokenFunc
	case "endfunc":
		return TokenEndFunc

	case "call":
		return TokenCall
	case "give":
		return TokenGive

	case "winput":
		return TokenInput
	case "wtime":
		return TokenTime
	case "wdate":
		return TokenDate
	case "wclear":
		return TokenClear
	case "wait":
		return TokenWait
	case "wexit":
		return TokenExit

	default:
		return TokenIdentifier
	}
}

func isIdentifierStart(ch byte) bool {
	return unicode.IsLetter(rune(ch)) || ch == '_'
}

func isIdentifierPart(ch byte) bool {
	return unicode.IsLetter(rune(ch)) ||
		unicode.IsDigit(rune(ch)) ||
		ch == '_'
}

func (l *Lexer) skipSpaces() {
	for !l.isAtEnd() {
		switch l.peek() {
		case ' ', '\t':
			l.advance()
		default:
			return
		}
	}
}

func (l *Lexer) advance() byte {
	if l.isAtEnd() {
		return 0
	}

	ch := l.source[l.position]
	l.position++

	if ch == '\n' {
		l.line++
		l.column = 1
	} else {
		l.column++
	}

	return ch
}

func (l *Lexer) peek() byte {
	if l.isAtEnd() {
		return 0
	}

	return l.source[l.position]
}

func (l *Lexer) isAtEnd() bool {
	return l.position >= len(l.source)
}

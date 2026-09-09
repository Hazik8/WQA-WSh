package parser

import (
	"fmt"

	"windroid/wqa/internal/lexer"
)

type Parser struct {
	tokens []lexer.Token
	pos    int
}

func New(tokens []lexer.Token) *Parser {
	return &Parser{
		tokens: tokens,
	}
}

func (p *Parser) Parse() (*Program, error) {
	program := &Program{}

	for !p.isAtEnd() {
		if p.match(lexer.TokenNewline) {
			continue
		}

		statement, err := p.parseStatement()
		if err != nil {
			return nil, err
		}

		program.Statements = append(
			program.Statements,
			statement,
		)

		for p.match(lexer.TokenNewline) {
		}
	}

	return program, nil
}

func (p *Parser) parseStatement() (Statement, error) {
	switch p.peek().Type {

	case lexer.TokenSet:
		return p.parseSet()

	case lexer.TokenPrint:
		return p.parsePrint()

	case lexer.TokenAdd:
		return p.parseBinary(lexer.TokenAdd)

	case lexer.TokenSub:
		return p.parseBinary(lexer.TokenSub)

	case lexer.TokenMul:
		return p.parseBinary(lexer.TokenMul)

	case lexer.TokenDiv:
		return p.parseBinary(lexer.TokenDiv)

	case lexer.TokenIf:
		return p.parseIf()

	default:
		token := p.peek()

		return nil, p.errorAt(
			token,
			fmt.Sprintf(
				"unexpected token %q",
				token.Lexeme,
			),
		)
	}
}

func (p *Parser) parseSet() (Statement, error) {
	keyword := p.advance()

	name, err := p.expect(
		lexer.TokenIdentifier,
		"expected variable name after wet",
	)
	if err != nil {
		return nil, err
	}

	if _, err := p.expect(
		lexer.TokenEquals,
		"expected '=' after variable name",
	); err != nil {
		return nil, err
	}

	value, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	_ = keyword

	return SetStatement{
		Name:  name,
		Value: value,
	}, nil
}

func (p *Parser) parsePrint() (Statement, error) {
	p.advance()

	value, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	return PrintStatement{
		Value: value,
	}, nil
}

func (p *Parser) parseBinary(
	operator lexer.TokenType,
) (Statement, error) {
	p.advance()

	left, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	right, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	result, err := p.expect(
		lexer.TokenIdentifier,
		"expected result variable",
	)
	if err != nil {
		return nil, err
	}

	return BinaryStatement{
		Operator: operator,
		Left:     left,
		Right:    right,
		Result:   result,
	}, nil
}

func (p *Parser) parseIf() (Statement, error) {
	p.advance()

	left, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	operator := p.peek()

	switch operator.Type {
	case lexer.TokenGreater,
		lexer.TokenLess,
		lexer.TokenEqualEqual:
		p.advance()

	default:
		return nil, p.errorAt(
			operator,
			"expected comparison operator after condition",
		)
	}

	right, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	condition := ComparisonExpression{
		Left:     left,
		Operator: operator.Type,
		Right:    right,
	}

	if p.match(lexer.TokenNewline) {
		for p.match(lexer.TokenNewline) {
		}
	}

	var thenStatements []Statement

	for !p.isAtEnd() &&
		p.peek().Type != lexer.TokenElse &&
		p.peek().Type != lexer.TokenEndIf {

		statement, err := p.parseStatement()
		if err != nil {
			return nil, err
		}

		thenStatements = append(
			thenStatements,
			statement,
		)

		for p.match(lexer.TokenNewline) {
		}
	}

	var elseStatements []Statement

	if p.match(lexer.TokenElse) {

		for p.match(lexer.TokenNewline) {
		}

		for !p.isAtEnd() &&
			p.peek().Type != lexer.TokenEndIf {

			statement, err := p.parseStatement()
			if err != nil {
				return nil, err
			}

			elseStatements = append(
				elseStatements,
				statement,
			)

			for p.match(lexer.TokenNewline) {
			}
		}
	}

	if _, err := p.expect(
		lexer.TokenEndIf,
		"expected endwif after wif",
	); err != nil {
		return nil, err
	}

	return IfStatement{
		Condition: condition,
		Then:      thenStatements,
		Else:      elseStatements,
	}, nil
}

func (p *Parser) parseExpression() (Expression, error) {
	token := p.peek()

	switch token.Type {

	case lexer.TokenIdentifier:
		p.advance()

		return IdentifierExpression{
			Token: token,
		}, nil

	case lexer.TokenNumber:
		p.advance()

		return NumberExpression{
			Token: token,
		}, nil

	case lexer.TokenString:
		p.advance()

		return StringExpression{
			Token: token,
		}, nil

	default:
		return nil, p.errorAt(
			token,
			fmt.Sprintf(
				"expected expression, got %q",
				token.Lexeme,
			),
		)
	}
}

func (p *Parser) expect(
	tokenType lexer.TokenType,
	message string,
) (lexer.Token, error) {
	if p.check(tokenType) {
		return p.advance(), nil
	}

	return lexer.Token{}, p.errorAt(
		p.peek(),
		message,
	)
}

func (p *Parser) match(
	tokenType lexer.TokenType,
) bool {
	if !p.check(tokenType) {
		return false
	}

	p.advance()
	return true
}

func (p *Parser) check(
	tokenType lexer.TokenType,
) bool {
	if p.isAtEnd() {
		return tokenType == lexer.TokenEOF
	}

	return p.peek().Type == tokenType
}

func (p *Parser) advance() lexer.Token {
	if !p.isAtEnd() {
		p.pos++
	}

	return p.previous()
}

func (p *Parser) peek() lexer.Token {
	if p.pos >= len(p.tokens) {
		return lexer.Token{
			Type: lexer.TokenEOF,
			Line: p.lastLine(),
		}
	}

	return p.tokens[p.pos]
}

func (p *Parser) previous() lexer.Token {
	if p.pos == 0 {
		return lexer.Token{
			Type: lexer.TokenEOF,
		}
	}

	return p.tokens[p.pos-1]
}

func (p *Parser) isAtEnd() bool {
	return p.peek().Type == lexer.TokenEOF
}

func (p *Parser) lastLine() int {
	if len(p.tokens) == 0 {
		return 1
	}

	return p.tokens[len(p.tokens)-1].Line
}

func (p *Parser) errorAt(
	token lexer.Token,
	message string,
) error {
	return fmt.Errorf(
		"parser error at %d:%d: %s",
		token.Line,
		token.Column,
		message,
	)
}

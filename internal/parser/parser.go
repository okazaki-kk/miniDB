// Package parser implements a parser for the NanoDB's SQL dialect.
package parser

import (
	"fmt"

	"github.com/okazaki-kk/miniDB/internal/parser/ast"
	"github.com/okazaki-kk/miniDB/internal/parser/lexer"
	"github.com/okazaki-kk/miniDB/internal/parser/token"
)

// Parser takes a Lexer and builds an abstract syntax tree.
type Parser struct {
	lexer     *lexer.Lexer
	token     token.Token
	peekToken token.Token
}

// New returns new Parser.
func New(lx *lexer.Lexer) *Parser {
	return &Parser{
		lexer:     lx,
		token:     lx.NextToken(),
		peekToken: lx.NextToken(),
	}
}

// Parse parses the sql and returns a statement.
func (p *Parser) Parse() (ast.Statement, error) {
	// For simplicity, we parse one statement at a time but in the next release,
	// we should implement parsing multiple statements separated semicolon.
	return p.parseStatement()
}

func (p *Parser) nextToken() {
	p.token = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) parseStatement() (ast.Statement, error) {
	switch p.token.Type {
	// DML
	case token.SELECT:
		return p.parseSelectStatement()
	case token.EOF:
		return nil, nil
	default:
		return nil, fmt.Errorf("unexpected statement: %s(%q)", p.token.Type, p.token.Literal)
	}
}

func (p *Parser) parseSelectStatement() (ast.Statement, error) {
	p.nextToken()

	result, err := p.parseResultStatement()
	if err != nil {
		return nil, err
	}

	from, err := p.parseFromStatement()
	if err != nil {
		return nil, err
	}

	where, err := p.parseWhereStatement()
	if err != nil {
		return nil, err
	}

	selectStmt := ast.SelectStatement{
		Result: result,
		From:   from,
		Where:  where,
	}

	return &selectStmt, nil
}

func (p *Parser) parseResultStatement() ([]ast.ResultStatement, error) {
	var results []ast.ResultStatement

	for p.token.Type != token.EOF && p.token.Type != token.FROM {
		result, err := p.parseResult()
		if err != nil {
			return nil, err
		}

		results = append(results, result)

		if p.token.Type == token.EOF || p.token.Type == token.FROM {
			break
		}

		p.nextToken()
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("no columns specified")
	}

	return results, nil
}

func (p *Parser) parseResult() (ast.ResultStatement, error) {
	var (
		result ast.ResultStatement
		err    error
	)

	result.Expr, err = p.parsePrimaryExpr()
	if err != nil {
		return ast.ResultStatement{}, err
	}

	return result, nil
}

func (p *Parser) parseFromStatement() (*ast.FromStatement, error) {
	if p.token.Type != token.FROM {
		return nil, fmt.Errorf("unexpected token %q, expected where", p.token.Type)
	}

	p.nextToken()

	table, err := p.parseIdent()
	if err != nil {
		return nil, err
	}

	from := ast.FromStatement{
		Table: table.Name,
	}

	return &from, nil
}

func (p *Parser) parseWhereStatement() (*ast.WhereStatement, error) {
	if p.token.Type != token.WHERE {
		return nil, nil
	}

	p.nextToken()

	expr, err := p.parsePrimaryExpr()
	if err != nil {
		return nil, err
	}

	p.nextToken()

	where := ast.WhereStatement{
		Expr: expr,
	}

	return &where, nil
}

func (p *Parser) parsePrimaryExpr() (ast.Expression, error) {
	expr, err := p.parseExpr()
	if err != nil {
		return nil, err
	}

	if p.peekToken.Type == token.COMMA {
		p.nextToken()
	}

	return expr, nil
}

func (p *Parser) parseExpr() (ast.Expression, error) {
	expr, err := p.parseOperand()
	if err != nil {
		return nil, err
	}

	for p.peekToken.Type != token.COMMA && p.peekToken.Type != token.SEMICOLON && p.peekToken.Type != token.EOF && p.peekToken.Type != token.FROM && p.peekToken.Type != token.WHERE {
		p.nextToken()

		expr, err = p.parseConditionExpr(expr)
		if err != nil {
			return nil, err
		}
	}

	return expr, nil
}

func (p *Parser) parseOperand() (ast.Expression, error) {
	switch p.token.Type {
	case token.IDENT:
		return &ast.IdentExpr{Name: p.token.Literal}, nil
	case token.ASTERISK:
		return &ast.AsteriskExpr{}, nil
	case token.INT, token.TEXT, token.TRUE, token.FALSE, token.NULL:
		return p.parseScalar(p.token.Type)
	case token.PLUS, token.MINUS:
		return nil, nil
	case token.LPAREN:
		return p.parseGroupExpr()
	default:
		return nil, fmt.Errorf("unexpected operand %q", p.token.Type)
	}
}

func (p *Parser) parseIdent() (*ast.IdentExpr, error) {
	if p.token.Type != token.IDENT {
		return nil, fmt.Errorf("unexpected token %q", p.token.Type)
	}

	ident := ast.IdentExpr{
		Name: p.token.Literal,
	}

	p.nextToken()

	return &ident, nil
}

func (p *Parser) parseScalar(expected token.TokenType) (ast.Expression, error) {
	if p.token.Type != expected {
		return nil, fmt.Errorf("unexpected scalar type %q", p.token.Type)
	}

	scalar := ast.ScalarExpr{
		Type:    p.token.Type,
		Literal: p.token.Literal,
	}

	return &scalar, nil
}

func (p *Parser) parseConditionExpr(left ast.Expression) (ast.Expression, error) {
	operator := p.token.Type
	p.nextToken()

	right, err := p.parseExpr()
	if err != nil {
		return nil, err
	}

	expr := ast.ConditionExpr{
		Left:     left,
		Operator: operator,
		Right:    right,
	}

	return &expr, nil
}

func (p *Parser) parseGroupExpr() (ast.Expression, error) {
	p.nextToken()

	expr, err := p.parseExpr()
	if err != nil {
		return nil, err
	}

	p.nextToken()

	if p.token.Type != token.RPAREN {
		return nil, fmt.Errorf("expected %q but found %q", token.RPAREN, p.token.Type)
	}

	return expr, nil
}

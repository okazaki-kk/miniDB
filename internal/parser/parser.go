// Package parser implements a parser for the NanoDB's SQL dialect.
package parser

import (
	"fmt"

	"github.com/okazaki-kk/miniDB/internal/parser/ast"
	"github.com/okazaki-kk/miniDB/internal/parser/lexer"
	"github.com/okazaki-kk/miniDB/internal/parser/token"
)

type Parser struct {
	lexer     *lexer.Lexer
	token     token.Token
	peekToken token.Token
}

func New(lx *lexer.Lexer) *Parser {
	return &Parser{
		lexer:     lx,
		token:     lx.NextToken(),
		peekToken: lx.NextToken(),
	}
}

func (p *Parser) Parse() (ast.Statement, error) {
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
	case token.INSERT:
		return p.parseInsertStatement()
	case token.CREATE:
		p.nextToken()
		return p.parseCreateStatement()
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

	order, err := p.parseOrderByStatement()
	if err != nil {
		return nil, err
	}

	limit, err := p.parseLimitStatement()
	if err != nil {
		return nil, err
	}

	offset, err := p.parseOffsetStatement()
	if err != nil {
		return nil, err
	}

	selectStmt := ast.SelectStatement{
		Result:  result,
		From:    from,
		Where:   where,
		OrderBy: order,
		Limit:   limit,
		Offset:  offset,
	}

	return &selectStmt, nil
}

func (p *Parser) parseInsertStatement() (ast.Statement, error) {
	p.nextToken()

	if err := p.expect(token.INTO); err != nil {
		return nil, err
	}

	table, err := p.parseIdent()
	if err != nil {
		return nil, err
	}

	columns, err := p.parseColumnsStatement()
	if err != nil {
		return nil, err
	}

	values, err := p.parseValuesStatement()
	if err != nil {
		return nil, err
	}

	insert := ast.InsertStatement{
		Table:   table.Name,
		Columns: columns,
		Values:  values,
	}

	return &insert, nil
}

func (p *Parser) parseCreateStatement() (ast.Statement, error) {
	if p.token.Type == token.DATABASE {
		return p.parseCreateDatabaseStatement()
	}

	p.nextToken()
	table, err := p.parseIdent()
	if err != nil {
		return nil, err
	}

	columns, err := p.parseColumns()
	if err != nil {
		return nil, err
	}

	create := ast.CreateTableStatement{
		Table:   table.Name,
		Columns: columns,
	}

	return &create, nil
}

func (p *Parser) parseCreateDatabaseStatement() (ast.Statement, error) {
	p.nextToken()

	database, err := p.parseIdent()
	if err != nil {
		return nil, err
	}

	create := ast.CreateDatabaseStatement{
		Database: database.Name,
	}

	return &create, nil
}

func (p *Parser) parseColumns() ([]ast.Column, error) {
	if p.token.Literal != "(" {
		return nil, fmt.Errorf("expected (, got %q", p.token.Literal)
	}

	p.nextToken()

	columns := make([]ast.Column, 0)

	for p.token.Type != token.EOF && p.token.Type != token.RPAREN {
		if p.token.Type == token.COMMA {
			p.nextToken()
		}

		column, err := p.parseColumn()
		if err != nil {
			return nil, err
		}

		columns = append(columns, column)
	}

	if p.token.Literal != ")" {
		return nil, fmt.Errorf("expected (, got %q", p.token.Literal)
	}
	p.nextToken()

	return columns, nil
}

func (p *Parser) parseColumn() (ast.Column, error) {
	columnName, err := p.parseIdent()
	if err != nil {
		return ast.Column{}, err
	}

	columnType, err := p.parseColumnType()
	if err != nil {
		return ast.Column{}, err
	}

	column := ast.Column{
		Name: columnName.Name,
		Type: columnType,
	}

	return column, nil
}

func (p *Parser) parseColumnType() (token.TokenType, error) {
	fmt.Println(p.token.Type, p.token.Literal)
	switch p.token.Type {
	case token.INT, token.TEXT, token.TRUE, token.FALSE:
		columnType := p.token.Type
		p.nextToken()

		return columnType, nil
	}

	return "", fmt.Errorf("unexpected column type: %q", p.token.Type)
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
		return nil, nil
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

func (p *Parser) parseOrderByStatement() (*ast.OrderByStatement, error) {
	if p.token.Type != token.ORDER {
		return nil, nil
	}

	p.nextToken()

	if err := p.expect(token.BY); err != nil {
		return nil, err
	}

	column, err := p.parseIdent()
	if err != nil {
		return nil, err
	}

	var direction token.TokenType = token.ASC

	switch p.token.Type {
	case token.ASC, token.DESC:
		direction = p.token.Type
		p.nextToken()
	}

	order := ast.OrderByStatement{
		Column:    column.Name,
		Direction: direction,
	}

	return &order, nil
}

func (p *Parser) parseLimitStatement() (*ast.LimitStatement, error) {
	if p.token.Type != token.LIMIT {
		return nil, nil
	}
	p.nextToken()

	value, err := p.parseScalar(token.INT)
	if err != nil {
		return nil, err
	}

	p.nextToken()

	limit := ast.LimitStatement{
		Value: value,
	}
	return &limit, nil
}

func (p *Parser) parseOffsetStatement() (*ast.OffsetStatement, error) {
	if p.token.Type != token.OFFSET {
		return nil, nil
	}
	p.nextToken()

	value, err := p.parseScalar(token.INT)
	if err != nil {
		return nil, err
	}

	p.nextToken()

	offset := ast.OffsetStatement{
		Value: value,
	}
	return &offset, nil
}

func (p *Parser) parsePrimaryExpr() (ast.Expression, error) {
	expr, err := p.parseExpr(LOWEST)
	if err != nil {
		return nil, err
	}

	if p.peekToken.Type == token.COMMA {
		p.nextToken()
	}

	return expr, nil
}

func (p *Parser) parseExpr(precedence int) (ast.Expression, error) {
	expr, err := p.parseOperand()
	if err != nil {
		return nil, err
	}

	for p.peekToken.Type != token.COMMA && precedence < p.peekPrecedence() {
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
	precedence := precedences[operator]

	p.nextToken()

	right, err := p.parseExpr(precedence)
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

	expr, err := p.parseExpr(LOWEST)
	if err != nil {
		return nil, err
	}

	p.nextToken()

	if p.token.Type != token.RPAREN {
		return nil, fmt.Errorf("expected %q but found %q", token.RPAREN, p.token.Type)
	}

	return expr, nil
}

func (p *Parser) parseColumnsStatement() ([]string, error) {
	var columns []string

	if err := p.expect(token.LPAREN); err != nil {
		return nil, err
	}

	for p.token.Type != token.EOF && p.token.Type != token.RPAREN {
		if p.token.Type == token.COMMA {
			p.nextToken()
		}

		column, err := p.parseIdent()
		if err != nil {
			return nil, err
		}

		columns = append(columns, column.Name)
	}

	if err := p.expect(token.RPAREN); err != nil {
		return nil, err
	}

	return columns, nil
}

func (p *Parser) parseValuesStatement() ([]ast.Expression, error) {
	var values []ast.Expression

	if err := p.expect(token.VALUES); err != nil {
		return nil, err
	}

	if err := p.expect(token.LPAREN); err != nil {
		return nil, err
	}

	for p.token.Type != token.EOF && p.token.Type != token.RPAREN {
		expr, err := p.parsePrimaryExpr()
		if err != nil {
			return nil, err
		}

		values = append(values, expr)

		p.nextToken()
	}

	if err := p.expect(token.RPAREN); err != nil {
		return nil, err
	}

	return values, nil
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) expect(tokenType token.TokenType) error {
	defer p.nextToken()

	if p.token.Type == tokenType {
		return nil
	}

	return fmt.Errorf(
		"expected %q but found %q (%s)",
		tokenType,
		p.token.Literal,
		p.token.Type,
	)
}

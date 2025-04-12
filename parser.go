package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// TokenType represents the different kinds of tokens in a formula
type TokenType int

const (
	TOKEN_NUMBER TokenType = iota
	TOKEN_STRING
	TOKEN_CELL_REF
	TOKEN_RANGE
	TOKEN_OPERATOR
	TOKEN_FUNCTION
	TOKEN_COMMA
	TOKEN_LPAREN
	TOKEN_RPAREN
	TOKEN_EOF
)

// Token represents a single element in a formula
type Token struct {
	Type    TokenType
	Value   string
	Literal interface{} // For numbers, this holds the actual float64
}

// Operator precedence levels
const (
	PREC_NONE     = 0
	PREC_ADD_SUB  = 1
	PREC_MUL_DIV  = 2
	PREC_UNARY    = 3
	PREC_FUNCTION = 4
)

// Tokenizer breaks a formula string into tokens
type Tokenizer struct {
	input   string
	start   int
	current int
	tokens  []Token
}

// Expression interface for all AST nodes
type Expression interface {
	Evaluate(s *sheet) (interface{}, error)
}

// LiteralExpr represents a numeric constant or string literal
type LiteralExpr struct {
	Value interface{} // float64 or string
}

// BinaryExpr represents operations like A1 + B2
type BinaryExpr struct {
	Left     Expression
	Operator string
	Right    Expression
}

// FunctionExpr represents function calls like SUM(A1:B3)
type FunctionExpr struct {
	Name string
	Args []Expression
}

// CellRefExpr represents a reference to a single cell
type CellRefExpr struct {
	Position vector
}

// RangeExpr represents a range of cells like A1:B3
type RangeExpr struct {
	Start vector
	End   vector
}

// UnaryExpr represents operations like -A1
type UnaryExpr struct {
	Operator string
	Right    Expression
}

// GroupingExpr represents parenthesized expressions
type GroupingExpr struct {
	Expression Expression
}

// NewTokenizer creates a new tokenizer for the given input string
func NewTokenizer(input string) *Tokenizer {
	// Remove the leading = if present
	if len(input) > 0 && input[0] == '=' {
		input = input[1:]
	}
	return &Tokenizer{
		input:   input,
		start:   0,
		current: 0,
		tokens:  []Token{},
	}
}

// Tokenize breaks the input string into tokens
func (t *Tokenizer) Tokenize() ([]Token, error) {
	for !t.isAtEnd() {
		t.start = t.current
		err := t.scanToken()
		if err != nil {
			return nil, err
		}
	}

	t.tokens = append(t.tokens, Token{Type: TOKEN_EOF, Value: ""})
	return t.tokens, nil
}

// Main token scanning method
func (t *Tokenizer) scanToken() error {
	c := t.advance()

	switch c {
	case '(':
		t.addToken(TOKEN_LPAREN, "(")
	case ')':
		t.addToken(TOKEN_RPAREN, ")")
	case '+':
		t.addToken(TOKEN_OPERATOR, "+")
	case '-':
		t.addToken(TOKEN_OPERATOR, "-")
	case '*':
		t.addToken(TOKEN_OPERATOR, "*")
	case '/':
		t.addToken(TOKEN_OPERATOR, "/")
	case ',':
		t.addToken(TOKEN_COMMA, ",")
	case ':':
		t.addToken(TOKEN_RANGE, ":")
	case '"':
		return t.string()
	case ' ', '\t', '\r', '\n':
		// Ignore whitespace
	default:
		if isDigit(c) {
			return t.number()
		} else if isAlpha(c) {
			return t.identifier()
		} else {
			return fmt.Errorf("unexpected character: %c", c)
		}
	}
	return nil
}

// Parse cell references like A1 or function names or operators
func (t *Tokenizer) identifier() error {
	for isAlphaNumeric(t.peek()) {
		t.advance()
	}

	value := t.input[t.start:t.current]
	
	// Check if it's a cell reference (like A1)
	if isCellReference(value) {
		t.addToken(TOKEN_CELL_REF, value)
		return nil
	}
	
	// Check if it's a function name
	if t.peek() == '(' {
		t.addToken(TOKEN_FUNCTION, strings.ToUpper(value))
		return nil
	}
	
	// Otherwise it's an invalid token
	return fmt.Errorf("invalid identifier: %s", value)
}

// Parse string literals
func (t *Tokenizer) string() error {
	for t.peek() != '"' && !t.isAtEnd() {
		t.advance()
	}

	if t.isAtEnd() {
		return errors.New("unterminated string")
	}

	// Consume the closing "
	t.advance()

	// Extract the string value (without the quotes)
	value := t.input[t.start+1 : t.current-1]
	t.addToken(TOKEN_STRING, value)
	return nil
}

// Parse numeric literals
func (t *Tokenizer) number() error {
	for isDigit(t.peek()) {
		t.advance()
	}

	// Look for a decimal part
	if t.peek() == '.' && isDigit(t.peekNext()) {
		// Consume the "."
		t.advance()

		// Consume the decimal digits
		for isDigit(t.peek()) {
			t.advance()
		}
	}

	value := t.input[t.start:t.current]
	number, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return fmt.Errorf("invalid number: %s", value)
	}
	
	token := Token{
		Type:    TOKEN_NUMBER,
		Value:   value,
		Literal: number,
	}
	t.tokens = append(t.tokens, token)
	return nil
}

// Parser parses tokens into an expression tree
type Parser struct {
	tokens  []Token
	current int
}

// NewParser creates a new parser for the given tokens
func NewParser(tokens []Token) *Parser {
	return &Parser{
		tokens:  tokens,
		current: 0,
	}
}

// Parse the tokens into an expression
func (p *Parser) Parse() (Expression, error) {
	return p.expression()
}

// Parse an expression
func (p *Parser) expression() (Expression, error) {
	return p.additive()
}

// Parse addition and subtraction
func (p *Parser) additive() (Expression, error) {
	expr, err := p.multiplicative()
	if err != nil {
		return nil, err
	}

	for p.match(TOKEN_OPERATOR) && (p.previous().Value == "+" || p.previous().Value == "-") {
		operator := p.previous().Value
		right, err := p.multiplicative()
		if err != nil {
			return nil, err
		}
		expr = &BinaryExpr{Left: expr, Operator: operator, Right: right}
	}

	return expr, nil
}

// Parse multiplication and division
func (p *Parser) multiplicative() (Expression, error) {
	expr, err := p.unary()
	if err != nil {
		return nil, err
	}

	for p.match(TOKEN_OPERATOR) && (p.previous().Value == "*" || p.previous().Value == "/") {
		operator := p.previous().Value
		right, err := p.unary()
		if err != nil {
			return nil, err
		}
		expr = &BinaryExpr{Left: expr, Operator: operator, Right: right}
	}

	return expr, nil
}

// Parse unary operators
func (p *Parser) unary() (Expression, error) {
	if p.match(TOKEN_OPERATOR) && (p.previous().Value == "-" || p.previous().Value == "+") {
		operator := p.previous().Value
		right, err := p.unary()
		if err != nil {
			return nil, err
		}
		return &UnaryExpr{Operator: operator, Right: right}, nil
	}

	return p.primary()
}

// Parse primary expressions (literals, references, ranges, functions)
func (p *Parser) primary() (Expression, error) {
	if p.match(TOKEN_NUMBER) {
		return &LiteralExpr{Value: p.previous().Literal}, nil
	}

	if p.match(TOKEN_STRING) {
		return &LiteralExpr{Value: p.previous().Value}, nil
	}

	if p.match(TOKEN_CELL_REF) {
		pos := alphaNumericToPosition(p.previous().Value)
		
		// Check if it's a range (A1:B3)
		if p.match(TOKEN_RANGE) {
			if !p.match(TOKEN_CELL_REF) {
				return nil, errors.New("expected cell reference after :")
			}
			endPos := alphaNumericToPosition(p.previous().Value)
			return &RangeExpr{Start: pos, End: endPos}, nil
		}
		
		return &CellRefExpr{Position: pos}, nil
	}

	if p.match(TOKEN_FUNCTION) {
		funcName := p.previous().Value

		if !p.match(TOKEN_LPAREN) {
			return nil, errors.New("expected '(' after function name")
		}

		args := []Expression{}
		
		// Check for empty argument list
		if !p.check(TOKEN_RPAREN) {
			// Parse arguments
			for {
				arg, err := p.expression()
				if err != nil {
					return nil, err
				}
				args = append(args, arg)

				if !p.match(TOKEN_COMMA) {
					break
				}
			}
		}

		if !p.match(TOKEN_RPAREN) {
			return nil, errors.New("expected ')' after function arguments")
		}

		return &FunctionExpr{Name: funcName, Args: args}, nil
	}

	if p.match(TOKEN_LPAREN) {
		expr, err := p.expression()
		if err != nil {
			return nil, err
		}

		if !p.match(TOKEN_RPAREN) {
			return nil, errors.New("expected ')' after expression")
		}

		return &GroupingExpr{Expression: expr}, nil
	}

	return nil, fmt.Errorf("unexpected token: %s", p.peek().Value)
}

// Helper methods for parser
func (p *Parser) match(types ...TokenType) bool {
	for _, typ := range types {
		if p.check(typ) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) check(typ TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().Type == typ
}

func (p *Parser) advance() Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}

func (p *Parser) isAtEnd() bool {
	return p.peek().Type == TOKEN_EOF
}

func (p *Parser) peek() Token {
	return p.tokens[p.current]
}

func (p *Parser) previous() Token {
	return p.tokens[p.current-1]
}

// Helper methods for tokenizer
func (t *Tokenizer) addToken(typ TokenType, value string) {
	t.tokens = append(t.tokens, Token{Type: typ, Value: value})
}

func (t *Tokenizer) advance() byte {
	t.current++
	return t.input[t.current-1]
}

func (t *Tokenizer) peek() byte {
	if t.isAtEnd() {
		return 0
	}
	return t.input[t.current]
}

func (t *Tokenizer) peekNext() byte {
	if t.current+1 >= len(t.input) {
		return 0
	}
	return t.input[t.current+1]
}

func (t *Tokenizer) isAtEnd() bool {
	return t.current >= len(t.input)
}

// Utility functions
func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		c == '_'
}

func isAlphaNumeric(c byte) bool {
	return isAlpha(c) || isDigit(c)
}

func isCellReference(s string) bool {
	if len(s) < 2 {
		return false
	}
	
	// Check if it starts with a letter followed by a number
	col := 0
	for col < len(s) && isAlpha(s[col]) {
		col++
	}
	
	if col == 0 || col == len(s) {
		return false
	}
	
	for i := col; i < len(s); i++ {
		if !isDigit(s[i]) {
			return false
		}
	}
	
	return true
}

// Evaluate methods for expression types
func (e *LiteralExpr) Evaluate(s *sheet) (interface{}, error) {
	return e.Value, nil
}

func (e *BinaryExpr) Evaluate(s *sheet) (interface{}, error) {
	left, err := e.Left.Evaluate(s)
	if err != nil {
		return nil, err
	}
	
	right, err := e.Right.Evaluate(s)
	if err != nil {
		return nil, err
	}
	
	// Try to convert both operands to float64
	leftNum, leftOk := toFloat(left)
	rightNum, rightOk := toFloat(right)
	
	if !leftOk || !rightOk {
		return nil, errors.New("#VALUE! - operands must be numbers")
	}
	
	switch e.Operator {
	case "+":
		return leftNum + rightNum, nil
	case "-":
		return leftNum - rightNum, nil
	case "*":
		return leftNum * rightNum, nil
	case "/":
		if rightNum == 0 {
			return nil, errors.New("#DIV/0! - division by zero")
		}
		return leftNum / rightNum, nil
	default:
		return nil, fmt.Errorf("unknown operator: %s", e.Operator)
	}
}

func (e *UnaryExpr) Evaluate(s *sheet) (interface{}, error) {
	right, err := e.Right.Evaluate(s)
	if err != nil {
		return nil, err
	}
	
	rightNum, ok := toFloat(right)
	if !ok {
		return nil, errors.New("#VALUE! - operand must be a number")
	}
	
	switch e.Operator {
	case "-":
		return -rightNum, nil
	case "+":
		return rightNum, nil
	default:
		return nil, fmt.Errorf("unknown unary operator: %s", e.Operator)
	}
}

func (e *GroupingExpr) Evaluate(s *sheet) (interface{}, error) {
	return e.Expression.Evaluate(s)
}

func (e *CellRefExpr) Evaluate(s *sheet) (interface{}, error) {
	// Check if the cell exists
	cell, exists := s.cells[e.Position]
	if !exists {
		return 0.0, nil // Empty cells evaluate to 0
	}
	
	// If the cell has a formula, we should get its computed value
	computed, exists := s.computed[e.Position]
	if exists {
		// Try to convert to a number if possible
		value, err := strconv.ParseFloat(computed, 64)
		if err == nil {
			return value, nil
		}
		return computed, nil // Return as string if not a number
	}
	
	// Otherwise, return the raw content
	return cell.content, nil
}

func (e *RangeExpr) Evaluate(s *sheet) (interface{}, error) {
	// Ranges can only be used in functions, not directly
	return nil, errors.New("#ERROR! - ranges can only be used in functions")
}

func (e *FunctionExpr) Evaluate(s *sheet) (interface{}, error) {
	// Collect arguments
	var args []interface{}
	for _, arg := range e.Args {
		// Special case for ranges

		if rangeExpr, isRange := arg.(*RangeExpr); isRange {

			minRow := int(min([]float64{float64(rangeExpr.Start.row), float64(rangeExpr.End.row)}))
			maxRow := int(max([]float64{float64(rangeExpr.Start.row), float64(rangeExpr.End.row)}))

			minCol := int(min([]float64{float64(rangeExpr.Start.col), float64(rangeExpr.End.col)}))
			maxCol := int(max([]float64{float64(rangeExpr.Start.col), float64(rangeExpr.End.col)}))

			// Expand the range and add all cells
			for row := minRow; row <= maxRow; row++ {
				for col := minCol; col <= maxCol; col++ {
					pos := vector{row: row, col: col}
					cellExpr := &CellRefExpr{Position: pos}
					val, err := cellExpr.Evaluate(s)
					if err != nil {
						return nil, err
					}
					args = append(args, val)
				}
			}
		} else {
			// Regular argument
			val, err := arg.Evaluate(s)
			if err != nil {
				return nil, err
			}
			args = append(args, val)
		}
	}
	
	// Convert all arguments to float64 if possible
	numbers := []float64{}
	for _, arg := range args {
		if num, ok := toFloat(arg); ok {
			numbers = append(numbers, num)
		}
	}
	
	// Evaluate the function
	switch e.Name {
	case "SUM":
		result := 0.0
		for _, num := range numbers {
			result += num
		}
		return result, nil
	case "PROD":
		if len(numbers) == 0 {
			return 0.0, nil
		}
		result := 1.0
		for _, num := range numbers {
			result *= num
		}
		return result, nil
	case "AVG":
		if len(numbers) == 0 {
			return 0.0, nil
		}
		sum := 0.0
		for _, num := range numbers {
			sum += num
		}
		return sum / float64(len(numbers)), nil
	case "MIN":
		if len(numbers) == 0 {
			return 0.0, nil
		}
		result := numbers[0]
		for _, num := range numbers {
			if num < result {
				result = num
			}
		}
		return result, nil
	case "MAX":
		if len(numbers) == 0 {
			return 0.0, nil
		}
		result := numbers[0]
		for _, num := range numbers {
			if num > result {
				result = num
			}
		}
		return result, nil
	case "COUNT":
		return float64(len(numbers)), nil
	default:
		return nil, fmt.Errorf("unknown function: %s", e.Name)
	}
}

// Utility function to convert interface{} to float64
func toFloat(value interface{}) (float64, bool) {
	switch v := value.(type) {
	case float64:
		return v, true
	case int:
		return float64(v), true
	case string:
		if num, err := strconv.ParseFloat(v, 64); err == nil {
			return num, true
		}
	}
	return 0, false
}

// ParseFormula is the main entry point for formula parsing
func ParseFormula(formula string) (Expression, error) {
	// Tokenize the formula
	tokenizer := NewTokenizer(formula)
	tokens, err := tokenizer.Tokenize()
	if err != nil {
		return nil, fmt.Errorf("tokenization error: %v", err)
	}
	
	// Parse the tokens
	parser := NewParser(tokens)
	expr, err := parser.Parse()
	if err != nil {
		return nil, fmt.Errorf("parsing error: %v", err)
	}
	
	return expr, nil
}

func (s *sheet) compute(content string) string {
	// If it doesn't start with =, it's not a formula
	if !strings.HasPrefix(content, "=") {
		return content
	}
	
	// Parse and evaluate the formula
	expr, err := ParseFormula(content)
	if err != nil {
		return fmt.Sprintf("#ERROR! %v", err)
	}
	
	result, err := expr.Evaluate(s)
	if err != nil {
		return fmt.Sprintf("#ERROR! %v", err)
	}
	
	// Format the result
	switch v := result.(type) {
	case float64:
		return fmt.Sprintf("%.*f", globalPrecisionLimit, v)
	case string:
		return v
	default:
		return fmt.Sprintf("%v", v)
	}
}
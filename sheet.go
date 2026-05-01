package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/expr-lang/expr"
)

var funcNameRe  = regexp.MustCompile(`\b([A-Za-z]+)\s*\(`)
var intLiteralRe = regexp.MustCompile(`\b\d+(\.\d+)?\b`)

type sheet struct {
	filename 	string
	saved 		bool
	size	 	vector
	cells     	map[vector]cell
	computed 	map[vector]string
	errors		map[vector]string
	cursor    	cursor
	selection 	[]vector
	history     []map[vector]cell
	viewport 	viewport
	evaluating	map[vector]bool
}

func (s *sheet) widestCellInCol(col int) int {
	widest := minColWidth
	for row := 1; row < s.size.row; row++ {
		v := vector{row: row, col: col}
		if thisWidth := utf8.RuneCountInString(s.computed[v]); thisWidth > widest {
			widest = thisWidth
		}
	}
	return widest
}

// recalculate computes the values for all cells in the sheet
func (s *sheet) recalculate() {
	s.errors = map[vector]string{}
	for position, cell := range s.cells {
		result, err := s.evaluate(cell.getRawContent())
		s.computed[position] = result
		if err != nil {
			s.errors[position] = err.Error()
		}
	}
}

// replaces all cell references in the expression with their computed values
func (s *sheet) replaceCellReferences(expr string) string {
	regex := regexp.MustCompile(`\b[A-Za-z]+\d+\b`) // e.g. "A1", "B2", etc.
	return regex.ReplaceAllStringFunc(expr, func(match string) string {

		v := alphaNumericToPosition(match)
		if s.evaluating[v] {
			return errorText
		}

		cell := s.cells[v]
		content := cell.getRawContent()
		if content == "" {
			return "0.0"
		}

		s.evaluating[v] = true
		val, _ := s.evaluate(content)
		delete(s.evaluating, v)
		return val
	})
}

// expand range references, then replace with computed values
func (s *sheet) rewriteExpression(expr string) string {
	expanded := expandRangeReferences(expr)
	replaced := s.replaceCellReferences(expanded)
	return replaced
}

func (s *sheet) evaluate(content string) (string, error) {
	if content == "" {
		return "", nil
	}
	if content[0] != '=' {
		return content, nil
	}

	exprBody := content[1:] // Remove the leading '=' sign
	rewritten := s.rewriteExpression(exprBody)
	rewritten = funcNameRe.ReplaceAllStringFunc(rewritten, func(m string) string {
		return strings.ToLower(m[:len(m)-1]) + "("
	})
	rewritten = intLiteralRe.ReplaceAllStringFunc(rewritten, func(m string) string {
		if strings.Contains(m, ".") {
			return m
		}
		return m + ".0"
	})
	rewritten = evaluateCustomFunctions(rewritten)
	result, err := expr.Eval(rewritten, nil)

	if err != nil {
		return errorText, err
	}
	if result == nil {
		return errorText, fmt.Errorf("nil result")
	}

	return fmt.Sprint(result), nil
}


// fetchReferencedCells scans the expression for individual cell references (e.g., "A1")
// and returns a map of the referenced cells.
func (s *sheet) fetchReferencedCells(str string) (map[vector]cell) {
	
	refcells := make(map[vector]cell)

	refs := extractReferences(str)
	positions := positionsFromReferences(refs)

	for _, position := range positions {
		refcells[position] = s.cells[position]
	}

	return refcells
}

func (s *sheet) clearCells() {
	s.cells = make(map[vector]cell)
	s.computed = make(map[vector]string)
}

func (s *sheet) saveForUndo() {
	cellsCopy := make(map[vector]cell, len(s.cells))
	for p, c := range s.cells {
		cellsCopy[p] = c
	}
	s.history = append(s.history, cellsCopy)
}

func (s *sheet) undo() {

	if len(s.history) == 1 {
		return
	}

	s.saved = false

	s.clearCells()
	s.history = s.history[:len(s.history)-1]
	s.cells = s.history[len(s.history)-1]
	s.recalculate()
}

func (s *sheet) save() {

	file, err := os.Create(s.filename)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    for pos, cell := range s.cells {
        line := fmt.Sprintf("%s%d@%s\n", columnToLetters(pos.col), pos.row, cell.content)
        _, err := file.WriteString(line)
        if err != nil {
            log.Fatal(err)
        }
    }

	s.saved = true
}

func (s *sheet) load() {

	file, err := os.OpenFile(s.filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	s.cells = map[vector]cell{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "@")
		s.cells[alphaNumericToPosition(parts[0])] = cell{content: parts[1]}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	s.recalculate()
	s.saveForUndo()
	s.saved = true
}
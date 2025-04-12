package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

type sheet struct {
	filename 	string
	saved 		bool
	size	 	vector
	cells     	map[vector]cell
	computed 	map[vector]string
	cursor    	cursor
	selection 	[]vector
	history     []map[vector]cell
	viewport 	viewport
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

func (s *sheet) calculate() {

	for position, cell := range s.cells {
		s.computed[position] = s.compute(cell.content)
	}

}

func (s *sheet) fetchReferencedCells(str string) (map[vector]cell) {
	
	refcells := make(map[vector]cell)

	refs := extractReferences(str)
	positions := positionsFromReferences(refs)

	for _, position := range positions {
		refcells[position] = s.cells[position]
	}

	return refcells
}

// func (s *sheet) compute(str string) string {

// 	operands := s.collectOperands(s.fetchReferencedCells(str))
// 	formula := strings.ToUpper(strings.Split(str, "(")[0])

// 	if len(operands) == 0 {
// 		return str
// 	}

// 	result := 0.00

// 	switch formula {
// 	case "=SUM":
// 		result = sum(operands)
// 	case "=PROD":
// 		result = product(operands)
// 	case "=MAX":
// 		result = max(operands)
// 	case "=MIN":
// 		result = min(operands)
// 	case "=AVG":
// 		result = average(operands)
// 	case "=COUNT":
// 		result = count(operands)
// 	}

// 	return fmt.Sprintf("%.*f", maxPrecision(operands), result)
// }

func (s *sheet) collectOperands(cells map[vector]cell) ([]float64) {

	operands := []float64{}

	for _, c := range cells {

		// ignore empty cells in calculations
		// otherwise, =PROD will always return 0
		// if there's an empty cell in the range
		if c.content == "" {
			continue
		}
		content := s.compute(c.content)
		value, _ := strconv.ParseFloat(content, 64)
		operands = append(operands, value)
	}

	return operands
}

func (s *sheet) clearCells() {
	s.cells = make(map[vector]cell)
	s.computed = make(map[vector]string)
}

// func (g *sheet) clearCellsAndHistory() {
// 	g.cells = make(map[vector]cell)
// 	g.computed = make(map[vector]string)
// 	g.history = []map[vector]cell{}
// }

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
	s.calculate()
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

	s.calculate()
	s.saveForUndo()
	s.saved = true
}
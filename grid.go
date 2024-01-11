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

type grid struct {
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

func (g *grid) widestCellInCol(col int) int {
	widest := minColWidth
	for row := 1; row < g.size.row; row++ {
		v := vector{row: row, col: col}
		if thisWidth := utf8.RuneCountInString(g.computed[v]); thisWidth > widest {
			widest = thisWidth
		}
	}
	return widest
}

func (g *grid) calculate() {

	for position, cell := range g.cells {
		g.computed[position] = g.compute(cell.content)
	}

}

func (g *grid) fetchReferencedCells(s string) (map[vector]cell) {
	
	refcells := make(map[vector]cell)

	refs := extractReferences(s)
	positions := positionsFromReferences(refs)

	for _, position := range positions {
		refcells[position] = g.cells[position]
	}

	return refcells
}

func (g *grid) compute(s string) string {

	operands := g.collectOperands(g.fetchReferencedCells(s))
	formula := strings.ToUpper(strings.Split(s, "(")[0])

	if len(operands) == 0 {
		return s
	}

	result := 0.00

	switch formula {
	case "=SUM":
		result = sum(operands)
	case "=PROD":
		result = product(operands)
	case "=MAX":
		result = max(operands)
	case "=MIN":
		result = min(operands)
	case "=AVG":
		result = average(operands)
	case "=COUNT":
		result = count(operands)
	}

	return fmt.Sprintf("%.*f", maxPrecision(operands), result)
}

func (g *grid) collectOperands(cells map[vector]cell) ([]float64) {

	operands := []float64{}

	for _, c := range cells {

		// ignore empty cells in calculations
		// otherwise, =PROD will always return 0
		// if there's an empty cell in the range
		if c.content == "" {
			continue
		}
		content := g.compute(c.content)
		value, _ := strconv.ParseFloat(content, 64)
		operands = append(operands, value)
	}

	return operands
}

func (g *grid) clearCells() {
	g.cells = make(map[vector]cell)
	g.computed = make(map[vector]string)
}

// func (g *grid) clearCellsAndHistory() {
// 	g.cells = make(map[vector]cell)
// 	g.computed = make(map[vector]string)
// 	g.history = []map[vector]cell{}
// }

func (g *grid) saveForUndo() {
	cellsCopy := make(map[vector]cell, len(g.cells))
	for p, c := range g.cells {
		cellsCopy[p] = c
	}
	g.history = append(g.history, cellsCopy)
}

func (g *grid) undo() {

	if len(g.history) == 1 {
		return
	}

	g.saved = false

	g.clearCells()
	g.history = g.history[:len(g.history)-1]
	g.cells = g.history[len(g.history)-1]
	g.calculate()
}

func (g *grid) save() {

	file, err := os.Create(g.filename)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    for pos, cell := range g.cells {
        line := fmt.Sprintf("%s%d@%s\n", columnToLetters(pos.col), pos.row, cell.content)
        _, err := file.WriteString(line)
        if err != nil {
            log.Fatal(err)
        }
    }

	g.saved = true
}

func (g *grid) load() {

	file, err := os.OpenFile(g.filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	g.cells = map[vector]cell{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "@")
		g.cells[alphaNumericToPosition(parts[0])] = cell{content: parts[1]}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	g.calculate()
	g.saveForUndo()
	g.saved = true
}
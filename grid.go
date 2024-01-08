package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Grid struct {
	filename 	string
	saved 		bool
	size	 	Vector
	cells     	map[Vector]Cell
	computed 	map[Vector]string
	cursor    	Cursor
	selection 	[]Vector
	history     []map[Vector]Cell
	viewport 	Viewport
}

func (g *Grid) WidestCell(col int) int {
	widest := minColWidth
	for row := 1; row < g.size.row; row++ {
		v := Vector{row: row, col: col}
		if len(g.computed[v]) > widest {
			widest = len(g.computed[v])
		}
	}
	return widest
}

func (g *Grid) cellFromString(s string) Cell {

	alphaPart, numericPart := splitAlphaNumeric(s)

	col := lettersToColumn(alphaPart)
	row, _ := strconv.Atoi(numericPart)

	return g.cells[Vector{row: row, col: col}]
}


func (g *Grid) Calculate() {

	for position, cell := range g.cells {
		g.computed[position] = g.Compute(cell.content)
	}

}

func (g *Grid) fetchReferencedCells(s string) (map[Vector]Cell) {
	
	refCells := make(map[Vector]Cell)

	refs := extractReferences(s)
	positions := positionsFromReferences(refs)

	for _, position := range positions {
		refCells[position] = g.cells[position]
	}

	return refCells
}

func (g *Grid) Compute(s string) string {

	operands := g.CollectOperands(g.fetchReferencedCells(s))
	formula := strings.ToUpper(strings.Split(s, "(")[0])

	if len(operands) == 0 {
		return s
	}

	result := 0.00

	switch formula {
	case "=SUM":
		result = Sum(operands)
	case "=PROD":
		result = Product(operands)
	case "=MAX":
		result = Max(operands)
	case "=MIN":
		result = Min(operands)
	case "=AVG":
		result = Average(operands)
	case "=COUNT":
		result = Count(operands)
	}

	return fmt.Sprintf("%.*f", maxPrecision(operands), result)
}

func (g *Grid) CollectOperands(cells map[Vector]Cell) ([]float64) {

	operands := []float64{}

	for _, c := range cells {
		content := g.Compute(c.content)
		value, _ := strconv.ParseFloat(content, 64)
		operands = append(operands, value)
	}

	return operands
}

func (g *Grid) ClearCells() {
	g.cells = make(map[Vector]Cell)
	g.computed = make(map[Vector]string)
}

func (g *Grid) ClearCellsAndHistory() {
	g.cells = make(map[Vector]Cell)
	g.computed = make(map[Vector]string)
	g.history = []map[Vector]Cell{}
}

func (g *Grid) SaveForUndo() {
	cellsCopy := make(map[Vector]Cell, len(g.cells))
	for p, c := range g.cells {
		cellsCopy[p] = c
	}
	g.history = append(g.history, cellsCopy)
}

func (g *Grid) Undo() {

	if len(g.history) == 1 {
		return
	}

	g.saved = false

	g.ClearCells()
	g.history = g.history[:len(g.history)-1]
	g.cells = g.history[len(g.history)-1]
	g.Calculate()
}

func (g *Grid) Save() {

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

func (g *Grid) Load() {

	file, err := os.OpenFile(g.filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	g.cells = map[Vector]Cell{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "@")
		g.cells[alphaNumericToPosition(parts[0])] = Cell{content: parts[1]}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	g.Calculate()
	g.SaveForUndo()
	g.saved = true
}
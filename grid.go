package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Grid struct {
	filename 	string
	saved 		bool
	size	 	VectorColRow // probably should rename as Vector
	cells     	map[VectorColRow]Cell
	computed 	map[VectorColRow]string
	cursor    	Cursor
	selection 	[]VectorColRow
	history     []map[VectorColRow]Cell
}

func (g *Grid) WidestCell(col int) int {
	widest := minColWidth
	for row := 1; row < g.size.row; row++ {
		v := VectorColRow{row: row, col: col}
		if len(g.computed[v]) > widest {
			widest = len(g.computed[v])
		}
	}
	return widest
}

func (g *Grid) CellFromString(s string) Cell {

	alphaPart, numericPart := splitAlphaNumeric(s)

	col := lettersToColumn(alphaPart)
	row, _ := strconv.Atoi(numericPart)

	return g.cells[VectorColRow{row: row, col: col}]
}

func (g *Grid) Calculate() {

	for position, cell := range g.cells {
		g.computed[position] = g.Compute(cell.content)
	}

}

func (g *Grid) FetchReferencedCells(s string) (map[VectorColRow]Cell) {
	
	refCells := make(map[VectorColRow]Cell)

	pattern := `=([A-Z]+)\(([A-Z]+)(\d+):([A-Z]+)(\d+)\)`
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(s)

	if matches == nil {
		return refCells // empty
	}

	startRow, _ := strconv.Atoi(matches[3]) // e.g. "5" -> 5
	startCol := lettersToColumn(matches[2]) // e.g. "B" -> 2
	endRow, _ := strconv.Atoi(matches[5])
	endCol := lettersToColumn(matches[4])

	for row := startRow; row <= endRow; row++ {
		for col := startCol; col <= endCol; col++ {
			v := VectorColRow{row: row, col: col}
			c := g.cells[v]
			refCells[v] = c
		}
	}

	return refCells
}

func (g *Grid) Compute(s string) string {

	operands := g.CollectOperands(g.FetchReferencedCells(s))

	if len(operands) == 0 {
		return s
	}

	result := 0.00

	// temporarily only sum
	result = Sum(operands)

	// switch matches[1] {
	// case "SUM":
	// 	result = Sum(operands)
	// case "PROD":
	// 	result = Product(operands)
	// case "MAX":
	// 	result = Max(operands)
	// case "MIN":
	// 	result = Min(operands)
	// case "AVG":
	// 	result = Average(operands)
	// case "COUNT":
	// 	result = Count(operands)
	// }

	return fmt.Sprintf("%.*f", maxPrecision(operands), result)
}

func (g *Grid) CollectOperands(cells map[VectorColRow]Cell) ([]float64) {

	operands := []float64{}

	for _, c := range cells {
		content := g.Compute(c.content)
		value, _ := strconv.ParseFloat(content, 64)
		operands = append(operands, value)
	}

	return operands
}

func (g *Grid) ClearCells() {
	g.cells = make(map[VectorColRow]Cell)
	g.computed = make(map[VectorColRow]string)
}

func (g *Grid) ClearCellsAndHistory() {
	g.cells = make(map[VectorColRow]Cell)
	g.computed = make(map[VectorColRow]string)
	g.history = []map[VectorColRow]Cell{}
}

func (g *Grid) SaveForUndo() {
	cellsCopy := make(map[VectorColRow]Cell, len(g.cells))
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

	g.cells = map[VectorColRow]Cell{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "@")
		g.cells[alphaNumericToVectorColRow(parts[0])] = Cell{content: parts[1]}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	g.Calculate()
	g.SaveForUndo()
	g.saved = true
}
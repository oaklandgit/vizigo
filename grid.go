package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Grid struct {
	filename 	string
	saved 		bool
	size	 	Position // probably should rename as Vector
	cells     	map[Position]Cell
	computed 	map[Position]string
	cursor    	Cursor
	selection 	[]Position
}

func (g *Grid) WidestCell(col int) int {
	widest := minColWidth
	for row := 1; row < g.size.row; row++ {
		p := Position{row: row, col: col}
		if len(g.computed[p]) > widest {
			widest = len(g.computed[p])
		}
	}
	return widest
}

func (g *Grid) CellFromString(s string) Cell {

	alphaPart, numericPart := splitAlphaNumeric(s)

	col := lettersToColumn(alphaPart)
	row, _ := strconv.Atoi(numericPart)

	return g.cells[Position{row: row, col: col}]
}

func (g *Grid) Calculate() {

	for position, cell := range g.cells {
		g.computed[position] = g.Compute(cell.content)
	}

}

func (g *Grid) Compute(s string) string {

	pattern := `=([A-Z]+)\(([A-Z]+)(\d+):([A-Z]+)(\d+)\)`
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(s)

	if matches == nil {
		return s
	}

	startRow, _ := strconv.Atoi(matches[3]) // e.g. "5" -> 5
	startCol := lettersToColumn(matches[2]) // e.g. "B" -> 2
	endRow, _ := strconv.Atoi(matches[5])
	endCol := lettersToColumn(matches[4])

	operands := collectOperands(g, startRow, startCol, endRow, endCol)

	result := 0.00

	switch matches[1] {
	case "SUM":
		result = Sum(operands)
	case "PROD":
		result = Product(operands)
	case "MAX":
		result = Max(operands)
	case "MIN":
		result = Min(operands)
	case "AVG":
		result = Average(operands)
	case "COUNT":
		result = Count(operands)
	}

	return fmt.Sprintf("%.*f", maxPrecision(operands), result)

	// return strconv.FormatFloat(result, 'f', -1, 64)

}

func collectOperands(g *Grid, startRow, startCol, endRow, endCol int) []float64 {
	// generate a list of values between startCell and endCell
	operands := []float64{}

	for row := startRow; row <= endRow; row++ {
		for col := startCol; col <= endCol; col++ {
			p := Position{row: row, col: col}
			content := g.Compute(p.GetCellContent(g))

			value, _ := strconv.ParseFloat(content, 64)
			operands = append(operands, value)
		}
	}

	return operands
}

func (g *Grid) Save() {

	file, err := os.Create(g.filename)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    for pos, cell := range g.cells {
        line := fmt.Sprintf("%s%d:%s\n", columnToLetters(pos.col), pos.row, cell.content)
        _, err := file.WriteString(line)
        if err != nil {
            log.Fatal(err)
        }
    }

	g.saved = true
}
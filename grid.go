package main

import "strconv"

type Grid struct {
	cells map[Position]Cell
	cursor Position
}

func (g *Grid) CellFromString(s string) Cell {

	alphaPart, numericPart := SplitAlphaNumeric(s)

	col := LettersToColumn(alphaPart)
	row, _ := strconv.Atoi(numericPart)

	return g.cells[Position{row: row, col: col}]
}
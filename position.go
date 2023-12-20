package main

import "strconv"

type Position struct {
	row, col int
}

func (p Position) ToString() string {
	return ColumnToLetters(p.col) + strconv.Itoa(p.row)
}
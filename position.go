package main

import "strconv"

type Position struct {
	row, col int
}

func (p Position) ToString() string {
	return columnToLetters(p.col) + strconv.Itoa(p.row)
}

func (p Position) GetCellContent(g *Grid) string {
	return g.cells[p].content
}

func (p Position) SetCellContent(g *Grid, content string) {
	g.cells[p] = Cell{content: content}
}

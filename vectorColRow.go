package main

import "strconv"

type VectorColRow struct {
	col, row int
}

func (v VectorColRow) ToString() string {
	return columnToLetters(v.col) + strconv.Itoa(v.row)
}

func (v VectorColRow) GetCellContent(g *Grid) string {
	return g.cells[v].content
}

func (v VectorColRow) SetCellContent(g *Grid, content string) {
	g.cells[v] = Cell{content: content}
	g.saved = false
}

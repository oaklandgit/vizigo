package main

import "strconv"

type Vector struct {
	col, row int
}

func (v Vector) ToString() string {
	return columnToLetters(v.col) + strconv.Itoa(v.row)
}

func (v Vector) GetCellContent(g *Grid, value bool) string {
		if value {
		return g.computed[v]
	} else {
		return g.cells[v].content
	}
}

func (v Vector) SetCellContent(g *Grid, content string) {
	g.cells[v] = Cell{content: content}
	g.saved = false
}

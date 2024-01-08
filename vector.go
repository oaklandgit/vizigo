package main

import "strconv"

type vector struct {
	col, row int
}

func (v vector) ToString() string {
	return columnToLetters(v.col) + strconv.Itoa(v.row)
}

func (v vector) GetcellContent(g *grid, value bool) string {
		if value {
		return g.computed[v]
	} else {
		return g.cells[v].content
	}
}

func (v vector) SetcellContent(g *grid, content string) {
	g.cells[v] = cell{content: content}
	g.saved = false
}

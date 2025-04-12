package main

import "strconv"

type vector struct {
	col, row int
}

func (v vector) toString() string {
	return columnToLetters(v.col) + strconv.Itoa(v.row)
}

func (v vector) getCellContent(s *sheet, value bool) string {
	if value {
		return s.computed[v]
	} else {
		return s.cells[v].content
	}
}

func (v vector) setCellContent(s *sheet, content string) {
	s.cells[v] = cell{content: content}
	s.saved = false
}

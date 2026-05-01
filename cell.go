package main

import (
	"fmt"
	"strconv"
)

type cell struct {
	content string
}

func (c *cell) getRawContent() string {
	return c.content
}

func (c *cell) render(s *sheet, v vector, referenced bool) string {

	fmtStr := ""

	if _, err := strconv.ParseFloat(s.computed[v], 64); err == nil {
		fmtStr = "%*s" // numeric, so right align
	} else {
		fmtStr = "%-*s" // not numeric, so left align
	}

	width := s.widestCellInCol(v.col)

	if s.cursor.vector == v {

		if s.cursor.editMode {
			// Clip to column width so content never overflows into adjacent cells.
			// Pad the plain content first so fmt.Sprintf measures correctly,
			// then apply underlineChar — ANSI codes must not be present during padding
			// or fmt.Sprintf counts escape bytes as visible characters and skips padding.
			content := []rune(c.content)
			editIdx := s.cursor.editIndex
			if len(content) > width {
				start := len(content) - width
				content = content[start:]
				editIdx -= start
			}
			plain := fmt.Sprintf(fmtStr, width, string(content))
			return cursorEditMode.Render(underlineChar(plain, editIdx))
		}
		return cursorSelected.Render(fmt.Sprintf(fmtStr, width, s.computed[v]))
	}
	
	if referenced {
		return cellReferenced.Render(fmt.Sprintf(fmtStr, width, s.computed[v]))
	}
	
	return cursorDeselected.Render(fmt.Sprintf(fmtStr, width, s.computed[v]))
	
}

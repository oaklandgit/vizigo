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
			return cursorEditMode.Render(fmt.Sprintf(fmtStr, width, underlineChar(c.content, s.cursor.editIndex)))
		}
		return cursorSelected.Render(fmt.Sprintf(fmtStr, width, s.computed[v]))
	}
	
	if referenced {
		return cellReferenced.Render(fmt.Sprintf(fmtStr, width, s.computed[v]))
	}
	
	return cursorDeselected.Render(fmt.Sprintf(fmtStr, width, s.computed[v]))
	
}

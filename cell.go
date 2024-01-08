package main

import (
	"fmt"
	"strconv"
)

type cell struct {
	content string
}

func (c *cell) render(g *grid, v vector, referenced bool) string {

	fmtStr := ""

	if _, err := strconv.ParseFloat(g.computed[v], 64); err == nil {
		fmtStr = "%*s" // numeric, so right align
	} else {
		fmtStr = "%-*s" // not numeric, so left align
	}

	width := g.widestCellInCol(v.col)

	if g.cursor.vector == v {

		if g.cursor.editMode {
			return cursorEditMode.Render(fmt.Sprintf(fmtStr, width, underlineChar(c.content, g.cursor.editIndex)))
		}
		return cursorSelected.Render(fmt.Sprintf(fmtStr, width, g.computed[v]))
	}
	
	if referenced {
		return cellReferenced.Render(fmt.Sprintf(fmtStr, width, g.computed[v]))
	}
	
	return cursorDeselected.Render(fmt.Sprintf(fmtStr, width, g.computed[v]))
	
}

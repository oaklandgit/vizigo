package main

import (
	"fmt"
	"strconv"
)

type Cell struct {
	content string
}

func (c *Cell) Render(g *Grid, v *VectorColRow, referenced bool) string {

	fmtStr := ""

	if _, err := strconv.ParseFloat(g.computed[*v], 64); err == nil {
		fmtStr = "%*s" // numeric, so right align
	} else {
		fmtStr = "%-*s" // not numeric, so left align
	}

	width := g.WidestCell(v.col)

	if g.cursor.position == *v {

		if g.cursor.editMode {
			return CursorEditMode.Render(fmt.Sprintf(fmtStr, width, underlineChar(c.content, g.cursor.editIndex)))
		}
		return CursorSelected.Render(fmt.Sprintf(fmtStr, width, g.computed[*v]))
	}
	
	if referenced {
		return CellReferenced.Render(fmt.Sprintf(fmtStr, width, g.computed[*v]))
	}
	
	return CursorDeselected.Render(fmt.Sprintf(fmtStr, width, g.computed[*v]))
	
}

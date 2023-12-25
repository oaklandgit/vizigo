package main

import (
	"fmt"
	"strconv"
)

type Cell struct {
	content string
}

func (c *Cell) Render(g *Grid, p *Position) string {

	fmtStr := ""

	if _, err := strconv.ParseFloat(g.computed[*p], 64); err == nil {
		fmtStr = "%*s" // numeric, so right align
	} else {
		fmtStr = "%-*s" // not numeric, so left align
	}

	width := g.WidestCell(p.col)

	if g.cursor.position == *p {

		if g.cursor.editMode {
			return CursorEditMode.Render(fmt.Sprintf(fmtStr, width, underlineChar(c.content, g.cursor.editIndex)))
		} else {
			return CursorSelected.Render(fmt.Sprintf(fmtStr, width, g.computed[*p]))
		}
	}
	return CursorDeselected.Render(fmt.Sprintf(fmtStr, width, g.computed[*p]))
}

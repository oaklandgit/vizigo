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

	// cursor at this cell
	if g.cursor.position == *p {

		

		// edit mode
		if g.cursor.editMode {
			return CursorEditMode.Render(fmt.Sprintf(fmtStr, width, UnderlineChar(c.content, g.cursor.editIndex)))
			// return CursorEditMode.Render(PadStringToCenter(UnderlineChar(c.content, g.cursor.editIndex), width))
		} else {
			return CursorSelected.Render(fmt.Sprintf(fmtStr, width, g.computed[*p]))

			// return CursorSelected.Render(PadStringToCenter(g.computed[*p], width))
		}
	}

	// return PadStringToCenter(g.computed[*p], width)
	return CursorDeselected.Render(fmt.Sprintf(fmtStr, width, g.computed[*p]))

}

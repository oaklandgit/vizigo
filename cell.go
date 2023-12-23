package main

import (
	"fmt"
	"strconv"
)

type Cell struct {
	content string
}

func (c *Cell) Render(g Grid, p Position) string {

	fmtStr := ""

	if _, err := strconv.ParseFloat(g.computed[p], 64); err == nil {
		fmtStr = "%*s" // numeric, so right align
	} else {
		fmtStr = "%-*s" // not numeric, so left align
	}

	// cursor at this cell
	if g.cursor.position == p {

		// edit mode
		if g.cursor.editMode {
			return CursorEditMode.Render(c.content)
		} else {
			return CursorSelected.Render(g.computed[p])
		}
	}

	return fmt.Sprintf(fmtStr, ColWidth, g.computed[p])

}

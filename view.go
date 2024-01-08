package main

import (
	"fmt"
)

func (g grid) View() string {

	modeString := ""
	returnString := ""
	referenced := g.fetchReferencedcells(g.cursor.vector.GetcellContent(&g, false))

	// Status Bar ////
	if g.cursor.editMode {
		modeString = "EDIT "
	}
	returnString += fmt.Sprintf("\n%s%s %s",
		modeString,
		g.cursor.vector.ToString(),
		g.cursor.vector.GetcellContent(&g, false),
	)

	// find the min of the viewport size and the grid size
	
	rowsToRender := g.viewport.offset.row + g.viewport.size.row
	colsToRender := g.viewport.offset.col + g.viewport.size.col

	// Header ////
	returnString += "\n" + fmt.Sprintf("%-*s", firstColWidth, " ")
	for col := g.viewport.offset.col; col < colsToRender; col++ {

		width := g.widestCell(col)

		if col == g.cursor.vector.col {
			returnString += ThSelected.Render(padStringToCenter(columnToLetters(col), width))
		} else {
			returnString += ThDeselected.Render(padStringToCenter(columnToLetters(col), width))
		}
	}

	// Rows ////



	for row := g.viewport.offset.row; row < rowsToRender; row++ {

		returnString += "\n"

		if row == g.cursor.vector.row {
			returnString += TrSelected.Render(fmt.Sprintf("%d", row))
		} else {
			returnString += TrDeselected.Render(fmt.Sprintf("%d", row))
		}

		// Columns ////
		for col := g.viewport.offset.col; col < colsToRender; col++ {

			// cell

			v := vector{col: col, row: row}
			cell := g.cells[v]

			_, isRef := referenced[v]

			if isRef {
				returnString += cell.Render(&g, v, true)
			} else {
				returnString += cell.Render(&g, v, false)
			}

		}
	}

	if (g.saved) {
		returnString += "\n\n" + g.filename + " (saved)"
	} else {
		returnString += "\n\n" + g.filename + " (unsaved)"
	}

	returnString += helpText

	return returnString
}

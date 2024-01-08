package main

import (
	"fmt"
)

func (g Grid) View() string {

	modeString := ""
	returnString := ""
	referenced := g.fetchReferencedCells(g.cursor.Vector.GetCellContent(&g))

	// Status Bar ////
	if g.cursor.editMode {
		modeString = "EDIT "
	}
	returnString += fmt.Sprintf("\n%s%s %s",
		modeString,
		g.cursor.Vector.ToString(),
		g.cursor.Vector.GetCellContent(&g),
	)

	// find the min of the viewport size and the grid size
	
	rowsToRender := g.viewport.offset.row + g.viewport.size.row
	colsToRender := g.viewport.offset.col + g.viewport.size.col

	// Header ////
	returnString += "\n" + fmt.Sprintf("%-*s", firstColWidth, " ")
	for col := g.viewport.offset.col; col < colsToRender; col++ {

		width := g.WidestCell(col)

		if col == g.cursor.Vector.col {
			returnString += ThSelected.Render(padStringToCenter(columnToLetters(col), width))
		} else {
			returnString += ThDeselected.Render(padStringToCenter(columnToLetters(col), width))
		}
	}

	// Rows ////



	for row := g.viewport.offset.row; row < rowsToRender; row++ {

		returnString += "\n"

		if row == g.cursor.Vector.row {
			returnString += TrSelected.Render(fmt.Sprintf("%d", row))
		} else {
			returnString += TrDeselected.Render(fmt.Sprintf("%d", row))
		}

		// Columns ////
		for col := g.viewport.offset.col; col < colsToRender; col++ {

			// Cell

			v := Vector{col: col, row: row}
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

	returnString += "\n\n==== HELP ====\n"

	for _, action := range helpTextKeys {
		returnString += fmt.Sprintf("\n%-6s %-8s", action, helpText[action])
	}

	return returnString
}

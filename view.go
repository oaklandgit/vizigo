package main

import (
	"fmt"
)

func (g grid) View() string {

	returnString := ""
	modeString := ""
	fileString := ""

	if g.cursor.editMode {
		modeString = "EDIT "
	}

	if g.saved {
		fileString = g.filename + " (saved)"
	} else {
		fileString = g.filename + " (unsaved)"
	}

	referenced := g.fetchReferencedCells(g.cursor.vector.getCellContent(&g, false))

	

	// Status Bar ////
	
	returnString += fmt.Sprintf("\n%-34s %s\n",
		modeString + " " + g.cursor.toString() + " " + g.cursor.getCellContent(&g, false),
		fileString,
	)

	// find the min of the viewport size and the grid size
	
	rowsToRender := g.viewport.offset.row + g.viewport.size.row
	colsToRender := g.viewport.offset.col + g.viewport.size.col

	// Header ////
	returnString += "\n" + fmt.Sprintf("%-*s", firstColWidth, " ")
	for col := g.viewport.offset.col; col < colsToRender; col++ {

		width := g.widestCellInCol(col)

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
				returnString += cell.render(&g, v, true)
			} else {
				returnString += cell.render(&g, v, false)
			}

		}
	}


	returnString += "\n\nhelp ⌃h\n"
	if showHelp {
		returnString += helpText
	}

	return returnString
}

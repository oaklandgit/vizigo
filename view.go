package main

import (
	"fmt"
)

func (g Grid) View() string {

	modeString := ""
	returnString := ""

	// Status Bar ////
	if g.cursor.editMode {
		modeString = "EDIT "
	}
	returnString += fmt.Sprintf("\n%s%s %s",
		modeString,
		g.cursor.position.ToString(),
		g.cursor.position.GetCellContent(&g),
	)

	// Header ////
	returnString += "\n" + fmt.Sprintf("%-*s", firstColWidth, " ")
	for col := hOffset; col < g.size.col + hOffset; col++ {

		width := g.WidestCell(col)

		if col == g.cursor.position.col {
			returnString += ThSelected.Render(padStringToCenter(columnToLetters(col), width))
		} else {
			returnString += ThDeselected.Render(padStringToCenter(columnToLetters(col), width))
		}
	}

	// Rows ////
	for row := vOffset; row < g.size.row + vOffset; row++ {

		returnString += "\n"

		if row == g.cursor.position.row {
			returnString += TrSelected.Render(fmt.Sprintf("%d", row))
		} else {
			returnString += TrDeselected.Render(fmt.Sprintf("%d", row))
		}

		// Columns ////
		for col := hOffset; col < g.size.col + hOffset; col++ {

			// Cell
			p := Position{row: row, col: col}
			cell := g.cells[p]
			returnString += cell.Render(&g, &p)

		}
	}

	if (g.saved) {
		returnString += "\n\n" + g.filename + ".viz (saved)"
	} else {
		returnString += "\n\n" + g.filename + ".viz (unsaved)"
	}
	
	returnString += "\n\n" + helpText
	return returnString
}

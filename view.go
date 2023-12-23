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
	returnString += "\n" + fmt.Sprintf("%-*s", FirstColWidth, " ")
	for col := HOffset; col < Cols; col++ {
		if col == g.cursor.position.col {
			returnString += ThSelected.Render(ColumnToLetters(col))
		} else {
			returnString += ThDeselected.Render(ColumnToLetters(col))
		}
	}

	// Rows ////
	for row := VOffset; row < Rows; row++ {

		returnString += "\n"

		if row == g.cursor.position.row {
			returnString += TrSelected.Render(fmt.Sprintf("%d", row))
		} else {
			returnString += TrDeselected.Render(fmt.Sprintf("%d", row))
		}

		// Columns ////
		for col := HOffset; col < Cols; col++ {

			// Cell
			p := Position{row: row, col: col}
			cell := g.cells[p]
			returnString += cell.Render(g, p)

		}
	}

	returnString += "\n\n" + HelpText
	return returnString
}

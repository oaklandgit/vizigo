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
	returnString += fmt.Sprintf("\n%s%s %s %d",
		modeString,
		g.cursor.position.ToString(),
		g.cursor.position.GetCellContent(&g),
		g.cursor.editIndex,
	)

	// Header ////
	returnString += "\n" + fmt.Sprintf("%-*s", FirstColWidth, " ")
	for col := HOffset; col < g.size.col; col++ {

		width := g.WidestCell(col)

		if col == g.cursor.position.col {
			returnString += ThSelected.Render(PadStringToCenter(ColumnToLetters(col), width))
		} else {
			returnString += ThDeselected.Render(PadStringToCenter(ColumnToLetters(col), width))
		}
	}

	// Rows ////
	for row := VOffset; row < g.size.row; row++ {

		returnString += "\n"

		if row == g.cursor.position.row {
			returnString += TrSelected.Render(fmt.Sprintf("%d", row))
		} else {
			returnString += TrDeselected.Render(fmt.Sprintf("%d", row))
		}

		// Columns ////
		for col := HOffset; col < g.size.col; col++ {

			// Cell
			p := Position{row: row, col: col}
			cell := g.cells[p]
			returnString += cell.Render(&g, &p)

		}
	}

	returnString += "\n\n" + HelpText
	return returnString
}

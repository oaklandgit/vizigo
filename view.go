package main

import (
	"fmt"
)

func (g Grid) View() string {

	modeString := ""
	returnString := ""
	referenced := g.FetchReferencedCells(g.cursor.position.GetCellContent(&g))



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

			v := VectorColRow{col: col, row: row}
			cell := g.cells[v]

			_, isRef := referenced[v]

			if isRef {
				returnString += cell.Render(&g, &v, true)
			} else {
				returnString += cell.Render(&g, &v, false)
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

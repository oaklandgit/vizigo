package main

import (
	"fmt"
)

func (s sheet) View() string {

	returnString := ""
	modeString := ""
	fileString := ""

	if s.cursor.editMode {
		modeString = "EDIT "
	}

	if s.saved {
		fileString = s.filename + " (saved)"
	} else {
		fileString = s.filename + " (unsaved)"
	}

	referenced := s.fetchReferencedCells(s.cursor.vector.getCellContent(&s, false))

	

	// Status Bar ////
	
	returnString += fmt.Sprintf("\n%-34s %s\n",
		modeString + " " + s.cursor.toString() + " " + s.cursor.getCellContent(&s, false),
		fileString,
	)

	// find the min of the viewport size and the sheet size
	
	rowsToRender := s.viewport.offset.row + s.viewport.size.row
	colsToRender := s.viewport.offset.col + s.viewport.size.col
	

	// rowsToRender := int(min([]float64{float64(rangeExpr.Start.row), float64(rangeExpr.End.row)}))
	// maxRow := int(max([]float64{float64(rangeExpr.Start.row), float64(rangeExpr.End.row)}))

	// Header ////
	returnString += "\n" + fmt.Sprintf("%-*s", firstColWidth, " ")
	for col := s.viewport.offset.col; col < colsToRender; col++ {

		width := s.widestCellInCol(col)

		if col == s.cursor.vector.col {
			returnString += ThSelected.Render(padStringToCenter(columnToLetters(col), width))
		} else {
			returnString += ThDeselected.Render(padStringToCenter(columnToLetters(col), width))
		}
	}

	// Rows ////
	for row := s.viewport.offset.row; row < rowsToRender; row++ {

		returnString += "\n"

		if row == s.cursor.vector.row {
			returnString += TrSelected.Render(fmt.Sprintf("%d", row))
		} else {
			returnString += TrDeselected.Render(fmt.Sprintf("%d", row))
		}

		// Columns ////
		for col := s.viewport.offset.col; col < colsToRender; col++ {

			// cell

			v := vector{col: col, row: row}
			cell := s.cells[v]

			_, isRef := referenced[v]

			if isRef {
				returnString += cell.render(&s, v, true)
			} else {
				returnString += cell.render(&s, v, false)
			}

		}
	}


	returnString += "\n\nhelp ⌃h\n"
	if showHelp {
		returnString += helpText
	}

	return returnString
}

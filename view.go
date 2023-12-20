package main

import (
	"fmt"
	"slices"
)

func (g Grid) View() string {
	s := ""
	cellContent := ""

	// status bar
	s += fmt.Sprintf("\n%s [%s]",
		g.cursor.ToString(),
		GetCellContent(g, g.cursor))

	// header
	s += "\n" + fmt.Sprintf("%-*s", FirstColWidth, " ")
	for col := HOffset; col < Cols; col++ {
		if col == g.cursor.col {
			s += ThSelected.Render(ColumnToLetters(col))
		} else {
			s += ThDeselected.Render(ColumnToLetters(col))
		}
	}

	// rows start at 1
	for row := VOffset; row < Rows; row++ {

		s += "\n"

		if row == g.cursor.row {
			s += TrSelected.Render(fmt.Sprintf("%d", row))
		} else {
			s += TrDeselected.Render(fmt.Sprintf("%d", row))
		}
		

		for col := HOffset; col < Cols; col++ {

			cellContent = ""

			for pos, cell := range g.cells {
				p := Position{row: row, col: col}
				if pos == p {
					cellContent = g.Compute(cell.content)
				}
			}

			p := Position{row: row, col: col}

			if p == g.cursor {
				s += CursorSelected.Render(cellContent)
			} else if slices.Contains(g.selection, p){
				s += Selected.Render(cellContent)
			} else {
				s += CursorDeselected.Render(cellContent)
			}

			cellContent = ""
		}
	}

	
	s += "\n\n" + HelpText

	return s
}
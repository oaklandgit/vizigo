package main

import (
	"fmt"
)

func (g Grid) View() string {
	s := ""
	cellContent := ""

	// status bar
	mode := ""
	if editMode {
		mode = "EDIT "
	}

	s += fmt.Sprintf("\n%s%s %s",
		mode,
		g.cursor.ToString(),
		GetCellContent(g, g.cursor),
	)

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

					raw := cell.content

					if p == g.cursor && editMode {
						cellContent = fmt.Sprintf("%-*s", 3, raw)
					} else {
						cellContent = fmt.Sprintf("%-*s", 3, g.Compute(raw))
					}
				}
			}

			p := Position{row: row, col: col}

			if p == g.cursor && !editMode {
				s += CursorSelected.Render(cellContent)
			} else if p == g.cursor && editMode {
				s += CursorEditMode.Render(cellContent)
			} else {
				s += CursorDeselected.Render(cellContent)
			}

			cellContent = ""
		}
	}

	
	s += "\n\n" + HelpText

	return s
}
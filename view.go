package main

import "fmt"

func (g Grid) View() string {
	s := ""
	cellContent := ""

	// status bar
	s += fmt.Sprintf("\n%s%d [%s]",
		IntToLetters(g.cursor.row),
		g.cursor.col,
		GetCellContent(g, g.cursor))

	// header
	s += "\n" + fmt.Sprintf("%-*s", FirstColWidth, " ")
	for col := HOffset; col < Cols; col++ {
		if col == g.cursor.col {
			s += ThSelected.Render(IntToLetters(col))
		} else {
			s += ThDeselected.Render(IntToLetters(col))
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
					cellContent = Solver(cell.content)
				}
			}

			p := Position{row: row, col: col}

			if g.cursor == p {
				s += CursorSelected.Render(cellContent)
			} else {
				s += CursorDeselected.Render(cellContent)
			}

			cellContent = ""
		}
	}

	
	s += "\n\n" + HelpText

	return s
}
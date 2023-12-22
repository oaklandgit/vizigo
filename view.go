package main

import (
	"fmt"
)

func (g Grid) View() string {

	mStr := "" // mode string
	rStr := "" // return string
	cStr := "" // content string
	p := Position{}

	// status bar
	if g.cursor.editMode {
		mStr = "EDIT "
	}

	rStr += fmt.Sprintf("\n%s%s %s",
		mStr,
		g.cursor.position.ToString(),
		g.cursor.position.GetCellContent(&g),
	)

	// header
	rStr += "\n" + fmt.Sprintf("%-*s", FirstColWidth, " ")
	for col := HOffset; col < Cols; col++ {
		if col == g.cursor.position.col {
			rStr += ThSelected.Render(ColumnToLetters(col))
		} else {
			rStr += ThDeselected.Render(ColumnToLetters(col))
		}
	}

	// rows start at 1
	for row := VOffset; row < Rows; row++ {

		rStr += "\n"

		if row == g.cursor.position.row {
			rStr += TrSelected.Render(fmt.Sprintf("%d", row))
		} else {
			rStr += TrDeselected.Render(fmt.Sprintf("%d", row))
		}
		

		for col := HOffset; col < Cols; col++ {

			cStr = ""
			p.row = row
			p.col = col

			for pos, cell := range g.cells {
				if pos == p {

					raw := cell.content

					if p == g.cursor.position && g.cursor.editMode {
						cStr = fmt.Sprintf("%-*s", ColWidth, raw)
					} else {
						cStr = fmt.Sprintf("%-*s", ColWidth, g.Compute(raw))
					}
				}
			}

			if p == g.cursor.position && !g.cursor.editMode {
				rStr += CursorSelected.Render(cStr)
			} else if p == g.cursor.position && g.cursor.editMode {
				rStr += CursorEditMode.Render(UnderlineChar(cStr, g.cursor.editIndex))
			} else {
				rStr += CursorDeselected.Render(cStr)
			}

			cStr = ""
		}
	}

	
	rStr += "\n\n" + HelpText

	return rStr
}
package main

import tea "github.com/charmbracelet/bubbletea"


func (g Grid) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			editMode = false
		case "enter":
			editMode = !editMode
		case "up":
			editMode = false
			if g.cursor.row > 1 {
				g.cursor.row--
			}
		case "down":
			editMode = false
			if g.cursor.row < Rows - 1 {
				g.cursor.row++
			}
		case "left":
			if !editMode && g.cursor.col > 1 {
				g.cursor.col--
			}
		case "right":
			if !editMode && g.cursor.col < Cols - 1 {
				g.cursor.col++
			}
		case "ctrl+c":
			editMode = false
			clipboard = GetCellContent(g, g.cursor)

		case "ctrl+v":
			editMode = false
			SetCellContent(&g, g.cursor, clipboard)

		case "backspace":
			if !editMode {
				SetCellContent(&g, g.cursor, "")
			} else {
				was := GetCellContent(g, g.cursor)
				if len(was) > 0 {
					SetCellContent(&g, g.cursor, was[:len(was)-1])
				}
			}
		case "ctrl+x":
			return g, tea.Quit
		default:
			if !editMode {
				return g, nil
			}
			was := GetCellContent(g, g.cursor)
			SetCellContent(&g, g.cursor, was + msg.String())
		
		}
		
	}

	return g, nil
}
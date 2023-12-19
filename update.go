package main

import tea "github.com/charmbracelet/bubbletea"

func (g Grid) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			if g.cursor.row > 1 {
				g.cursor.row--
			}
		case "down":
			if g.cursor.row < Rows - 1 {
				g.cursor.row++
			}
		case "left":
			if g.cursor.col > 1 {
				g.cursor.col--
			}
		case "right":
			if g.cursor.col < Cols - 1 {
				g.cursor.col++
			}
		case "ctrl+c":
			clipboard = GetCellContent(g, g.cursor)

		case "ctrl+v":
			SetCellContent(&g, g.cursor, clipboard)
			
		case "ctrl+x":
			return g, tea.Quit
		}
		
	}

	return g, nil
}
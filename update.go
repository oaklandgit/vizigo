package main

import tea "github.com/charmbracelet/bubbletea"


func (g Grid) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			g.cursor.editMode = false
			g.cursor.editIndex = 0
		case "enter":
			g.cursor.ToggleEditMode()
		case "up":
			g.cursor.Up()
		case "down":
			g.cursor.Down()
		case "left":
			g.cursor.Left()
		case "right":
			g.cursor.Right()
		case "ctrl+c":
			g.cursor.Copy(&g)
		case "ctrl+v":
			g.cursor.Paste(&g)

		case "backspace":
			g.cursor.Backspace(&g)
		case "ctrl+x":
			return g, tea.Quit
		default:
			g.cursor.Entry(&g, msg.String())		
		}
		
	}

	return g, nil
}
package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (g grid) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			g.cursor.Escape()
		case "enter":
			g.cursor.Enter(&g)
		case "up":
			g.cursor.Up()
		case "down":
			g.cursor.Down(&g)
		case "left":
			g.cursor.Left(&g)
		case "right":
			g.cursor.Right(&g)
		case "tab":
			g.cursor.Tab(&g)
		case "ctrl+c":
			g.cursor.Copy(&g)
		case "ctrl+n":
			g.cursor.CopyValue(&g)
		case "ctrl+v":
			g.cursor.Paste(&g)
		case "backspace":
			g.cursor.Backspace(&g)
		case "ctrl+z":
			g.Undo()
		case "ctrl+s":
			g.Save()
		case "ctrl+q":
			return g, tea.Quit
		default:
			g.cursor.TextEntry(&g, msg.String())
		}

	}

	g.calculate()
	return g, nil
}

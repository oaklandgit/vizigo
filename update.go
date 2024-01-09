package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (g grid) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+h":
			showHelp = !showHelp
		case "esc":
			g.cursor.escape()
		case "enter":
			g.cursor.enter(&g)
		case "up":
			g.cursor.up(&g)
		case "down":
			g.cursor.down(&g)
		case "left":
			g.cursor.left(&g)
		case "right":
			g.cursor.right(&g)
		case "tab":
			g.cursor.tab(&g)
		case "ctrl+c":
			g.cursor.copy(&g)
		case "ctrl+n":
			g.cursor.copyValue(&g)
		case "ctrl+v":
			g.cursor.paste(&g)
		case "backspace":
			g.cursor.backspace(&g)
		case "ctrl+z":
			g.undo()
		case "ctrl+s":
			g.save()
		case "ctrl+q":
			return g, tea.Quit
		default:
			g.cursor.textEntry(&g, msg.String())
		}

	}

	g.calculate()
	return g, nil
}

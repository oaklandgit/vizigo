package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (s sheet) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+h":
			showHelp = !showHelp
		case "esc":
			s.cursor.escape()
		case "enter":
			s.cursor.enter(&s)
		case "up":
			s.cursor.up(&s)
		case "down":
			s.cursor.down(&s)
		case "left":
			s.cursor.left(&s)
		case "right":
			s.cursor.right(&s)
		case "tab":
			s.cursor.tab(&s)
		case "ctrl+c":
			s.cursor.copy(&s)
		case "ctrl+n":
			s.cursor.copyValue(&s)
		case "ctrl+v":
			s.cursor.paste(&s)
		case "backspace":
			s.cursor.backspace(&s)
		case "ctrl+z":
			s.undo()
		case "ctrl+s":
			s.save()
		case "ctrl+q":
			return s, tea.Quit
		default:
			s.cursor.textEntry(&s, msg.String())
		}

	}

	s.recalculate()
	return s, nil
}

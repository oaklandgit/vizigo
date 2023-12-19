package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

var clipboard string

func main() {
	p := tea.NewProgram(initialGrid())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error bro: %v", err)
		os.Exit(1)
	}
}

func initialGrid() Grid {

	return Grid{
		cells: map[Position]Cell{
			{row: 1, col: 1}: {content: "Hello"},
			{row: 1, col: 2}: {content: "Goodbye"},
			{row: 4, col: 4}: {content: "Monday"},
		},
		cursor: Position{row: 1, col: 1},
	}
}

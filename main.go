package main

import (
	"fmt"
	"os"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {


	p := tea.NewProgram(initialGrid())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error bro: %v", err)
		os.Exit(1)
	}
}

func initialGrid() Grid {

	args := os.Args[1:]
	cols := defaultCols
	rows := defaultRows

	if len(args) == 2 {
		arg1, colErr := strconv.Atoi(args[0])
		arg2, rowErr := strconv.Atoi(args[1])

		if colErr == nil && rowErr == nil {
			cols = arg1
			rows = arg2
		}
	}

	return Grid{
		size:      Position{row: rows, col: cols},
		cells:     map[Position]Cell{},
		computed:  map[Position]string{},
		cursor:    Cursor{Position{row: 1, col: 1}, false, -1, ""},
		selection: []Position{},
	}
}

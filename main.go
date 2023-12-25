package main

import (
	"flag"
	"fmt"
	"os"

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

	cols := flag.Int("c", defaultCols, "The number of columns")
	rows := flag.Int("r", defaultRows, "The number of rows")
	filename := flag.String("f", "untitled", "The filename to open")
    

	flag.Parse()

	return Grid{
		filename:  *filename,
		saved:     false,
		size:      Position{row: *rows, col: *cols},
		cells:     map[Position]Cell{},
		computed:  map[Position]string{},
		cursor:    Cursor{Position{row: 1, col: 1}, false, -1, ""},
		selection: []Position{},
	}
}

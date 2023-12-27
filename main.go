package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {


	app := tea.NewProgram(initialGrid())
	if _, err := app.Run(); err != nil {
		fmt.Printf("Error bro: %v", err)
		os.Exit(1)
	}
}

func initialGrid() Grid {

	cols := flag.Int("c", defaultCols, "The number of columns")
	rows := flag.Int("r", defaultRows, "The number of rows")
	filename := flag.String("f", defaultFilename + fileExtension, "The filename to open")
	flag.Parse()

	// Get the actual file extension
	ext := filepath.Ext(*filename)

	// Compare the expected and actual extensions
	if ext == "" {
		*filename = *filename + fileExtension
	}

	g := Grid{
		filename:  *filename,
		saved:     false,
		size:      VectorColRow{col: *cols, row: *rows},
		cells:     map[VectorColRow]Cell{},
		computed:  map[VectorColRow]string{},
		cursor:    Cursor{VectorColRow{col: 1, row: 1}, false, -1, ""},
		selection: []VectorColRow{},
		history:   []map[VectorColRow]Cell{},
	}

	g.Load()

	return g
}

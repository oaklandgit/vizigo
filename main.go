package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {

	app := tea.NewProgram(initialgrid())
	if _, err := app.Run(); err != nil {
		fmt.Printf("Error bro: %v", err)
		os.Exit(1)
	}
}

func initialgrid() grid {

	cols := flag.Int("c", defaultCols, "The number of columns")
	rows := flag.Int("r", defaultRows, "The number of rows")
	vcols := flag.Int("vc", viewportCols, "The number of columns in the viewport")
	vrows := flag.Int("vr", viewportRows, "The number of rows in the viewport")
	filename := flag.String("f", defaultFilename + fileExtension, "The filename to open")
	flag.Parse()

	// Get the actual file extension
	ext := filepath.Ext(*filename)

	// Compare the expected and actual extensions
	if ext == "" {
		*filename = *filename + fileExtension
	}

	g := grid{
		filename:  *filename,
		saved:     false,
		size:      vector{col: *cols, row: *rows},
		cells:     map[vector]cell{},
		computed:  map[vector]string{},
		cursor:    cursor{vector{col: 1, row: 1}, false, -1, ""},
		selection: []vector{},
		history:   []map[vector]cell{},
		viewport:  viewport{vector{col: *vcols, row: *vrows}, vector{col: 1, row: 1}},
	}

	g.load()

	return g
}

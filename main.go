package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var deselected = lipgloss.NewStyle().
    Foreground(lipgloss.Color("#000000")).
    Background(lipgloss.Color("#ffffff")).
	Width(CellWidth).
	PaddingLeft(1)

var selected = lipgloss.NewStyle().
    Foreground(lipgloss.Color("#ffffff")).
    Background(lipgloss.Color("#000000")).
	Width(CellWidth).
	PaddingLeft(1)

const (
	Rows      = 10
	Cols      = 6
	CellWidth = 12
)

type cell struct {
	content  string
	row, col int
}

type grid struct {
	cells    []cell
	row, col int
}

func main() {
	p := tea.NewProgram(initialGrid())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error bro: %v", err)
		os.Exit(1)
	}
}

func initialGrid() grid {
	return grid{
		row: 0,
		col: 0,
		cells: []cell{
			{content: "hello", row: 0, col: 0},
			{content: "goodbye", row: 0, col: 1},
			{content: "morning", row: 0, col: 2},
			{content: "noon", row: 0, col: 3},
		},
	}
}

func (g grid) Init() tea.Cmd {
	return nil
}

func (g grid) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			if g.row > 0 {
				g.row--
			}
		case "down":
			if g.row < Rows {
				g.row++
			}
		case "left":
			if g.col > 0 {
				g.col--
			}
		case "right":
			if g.col < Cols {
				g.col++
			}
		case "q":
			return g, tea.Quit
		}
	}

	return g, nil
}

func (g grid) View() string {
	s := ""
	cellContent := ""
	for row := 0; row < Rows; row++ {
		s += "\n"
		for col := 0; col < Cols; col++ {

			cellContent = ""
			for _, cell := range g.cells {
				if row == cell.row && col == cell.col {
					cellContent = cell.content
				}
			}

			if row == g.row && col == g.col {
				s += selected.Render(cellContent)
			} else {
				s += deselected.Render(cellContent)
			}

			cellContent = ""
		}
	}

	return s
}

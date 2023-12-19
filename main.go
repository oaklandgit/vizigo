package main

import (
	"fmt"
	"os"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var clipboard string

var selected = lipgloss.NewStyle().
    Foreground(lipgloss.Color("#000000")).
    Background(lipgloss.Color("#FFFFFF")).
	Width(CellWidth).
	PaddingLeft(1)

var deselected = lipgloss.NewStyle().
    Foreground(lipgloss.Color("#FFFFFF")).
    Background(lipgloss.Color("#000000")).
	Width(CellWidth).
	PaddingLeft(1)

var blue = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#0000FF"))

var green = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#00FF00"))

var red = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#FF0000"))

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
			{content: "123", row: 1, col: 0},
			{content: "456", row: 2, col: 0},
			{content: "=sum(A1:A2)", row: 3, col: 0},
		},
	}
}

func (g grid) Init() tea.Cmd {
	return nil
}

// func sameCell(a, b cell) bool {
// 	return a.row == b.row && a.col == b.col
// }

// func setCell(g grid, row, col int, content string) {
// 	for _, c := range g.cells {
// 		if row == c.row && col == c.col {
// 			c.content = content
// 		}
// 	}

// 	return g
// }

// func getCell(row, col int, cells *[]cell) *cell {
// 	for _, cell := range cells {
// 		if row == cell.row && col == cell.col {
// 			return &cell
// 		}
// 	}
// 	return &cell{}
// }

func getCellContent(row, col int, cells []cell) string {
	for _, cell := range cells {
		if row == cell.row && col == cell.col {
			return cell.content
		}
	}
	return ""
}

func setCellContent(g *grid, row int, col int, content string) {
	// replace if cell already exists
	for i, cell := range g.cells {
		if row == cell.row && col == cell.col {
			g.cells[i].content = content
			return
		}
	}

	// otherwise append
	g.cells = append(g.cells, cell{content: content, row: row, col: col})

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
		case "ctrl+c":
			clipboard = getCellContent(g.row, g.col, g.cells)

		case "ctrl+v":
			setCellContent(&g, g.row, g.col, clipboard)
			
		case "q":
			return g, tea.Quit
		}
	}

	return g, nil
}

func solve(s string) string {

	// formula
	if s[0] == '=' {
		return red.Render(s)
	}

	// number
	if _, err := strconv.Atoi(s); err == nil {
		return green.Render(s)
	}

	// string
	return blue.Render(s)
	
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
					cellContent = solve(cell.content)
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

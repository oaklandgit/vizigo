package main

import (
	"fmt"
	"os"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	Rows      = 10
	Cols      = 6
	CellWidth = 12
	Orange = "#fcb826"
	Gray = "#363535"
)

var clipboard string

var thDeselected = lipgloss.NewStyle().
    Foreground(lipgloss.Color(Gray)).
    Background(lipgloss.Color(Orange)).
	Width(CellWidth).
	Align(lipgloss.Center)

var thSelected = lipgloss.NewStyle().
    Foreground(lipgloss.Color(Orange)).
    Background(lipgloss.Color(Gray)).
	Width(CellWidth).
	Align(lipgloss.Center)

var trDeselected = lipgloss.NewStyle().
    Foreground(lipgloss.Color(Gray)).
    Background(lipgloss.Color(Orange)).
	Width(5).
	Align(lipgloss.Center)

var trSelected = lipgloss.NewStyle().
    Foreground(lipgloss.Color(Orange)).
    Background(lipgloss.Color(Gray)).
	Width(5).
	Align(lipgloss.Center)

var cursorSelected = lipgloss.NewStyle().
    Foreground(lipgloss.Color("#000000")).
    Background(lipgloss.Color(Orange)).
	Width(CellWidth).
	PaddingLeft(1)

var cursorDeselected = lipgloss.NewStyle().
	Width(CellWidth).
	PaddingLeft(1)

var blue = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#0000FF"))

var green = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#00FF00"))

var red = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#FF0000"))

type position struct {
	row, col int
}

type cell struct {
	content  string
}

func main() {
	p := tea.NewProgram(initialGrid())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error bro: %v", err)
		os.Exit(1)
	}
}


type grid struct {
	cells map[position]cell
	cursor position
}

func initialGrid() grid {

	return grid{
		cells: map[position]cell{
			{row: 1, col: 1}: {content: "Hello"},
			{row: 1, col: 2}: {content: "Goodbye"},
			{row: 4, col: 4}: {content: "Monday"},
		},
		cursor: position{row: 1, col: 1},
	}
}

func (g grid) Init() tea.Cmd {
	return nil
}

func getCellContent(g grid, p position) string {
	for pos, cell := range g.cells {
		if pos == p {
			return cell.content
		}
	}
	return ""
}

func setCellContent(g *grid, p position, content string) {
	g.cells[p] = cell{content: content}
}

func (g grid) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			if g.cursor.row > 1 {
				g.cursor.row--
			}
		case "down":
			if g.cursor.row < Rows - 1 {
				g.cursor.row++
			}
		case "left":
			if g.cursor.col > 1 {
				g.cursor.col--
			}
		case "right":
			if g.cursor.col < Cols - 1 {
				g.cursor.col++
			}
		case "ctrl+c":
			clipboard = getCellContent(g, g.cursor)

		case "ctrl+v":
			setCellContent(&g, g.cursor, clipboard)
			
		case "ctrl+x":
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
	alpha := "_ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// header
	s += "\n" + trDeselected.Render("")
	for col := 1; col < Cols; col++ {
		if col == g.cursor.col {
			s += thSelected.Render(string(alpha[col]))
		} else {
			s += thDeselected.Render(string(alpha[col]))
		}
	}

	// rows start at 1
	for row := 1; row < Rows; row++ {

		// newline
		s += "\n"

		// row number
		if row == g.cursor.row {
			s += trSelected.Render(fmt.Sprintf("%d", row))
		} else {
			s += trDeselected.Render(fmt.Sprintf("%d", row))
		}
		

		for col := 1; col < Cols; col++ {

			cellContent = ""

			for pos, cell := range g.cells {
				p := position{row: row, col: col}
				if pos == p {
					cellContent = solve(cell.content)
				}
			}

			p := position{row: row, col: col}

			if g.cursor == p {
				s += cursorSelected.Render(cellContent)
			} else {
				s += cursorDeselected.Render(cellContent)
			}

			cellContent = ""
		}
	}

	s += "\n\nmove: → ← ↑ ↓, copy: ⌃c, paste: ⌃v, save: ⌃s, exit: ⌃x\n\n"
	// s += fmt.Sprintln(g)

	return s
}

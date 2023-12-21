package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

var clipboard string
var editMode bool = false

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
			{row: 1, col: 1}: {content: "Operand1:"},
			{row: 2, col: 1}: {content: "Operand2:"},
			{row: 3, col: 1}: {content: "Operand3:"},

			{row: 4, col: 1}: {content: "Sum:"},
			{row: 5, col: 1}: {content: "Product:"},
			{row: 6, col: 1}: {content: "Max:"},
			{row: 7, col: 1}: {content: "Min:"},
			{row: 8, col: 1}: {content: "Average:"},
			{row: 9, col: 1}: {content: "Count:"},

			{row: 1, col: 2}: {content: "100"},
			{row: 2, col: 2}: {content: "200.5"},
			{row: 3, col: 2}: {content: "300.5"},

			{row: 4, col: 2}: {content: "=SUM(B1:B3)"},
			{row: 5, col: 2}: {content: "=PRODUCT(B1:B3)"},
			{row: 6, col: 2}: {content: "=MAX(B1:B3)"},
			{row: 7, col: 2}: {content: "=MIN(B1:B3)"},
			{row: 8, col: 2}: {content: "=AVERAGE(B1:B3)"},
			{row: 9, col: 2}: {content: "=COUNT(B1:B3)"},
		},
		cursor: Position{row: 1, col: 1},
		selection: []Position{},
	}
}

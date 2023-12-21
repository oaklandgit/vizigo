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
			{row: 1, col: 1}: {content: "Operand1:"},
			{row: 2, col: 1}: {content: "Operand2:"},
			{row: 3, col: 1}: {content: "---------"},

			{row: 4, col: 1}: {content: "Sum:"},
			{row: 5, col: 1}: {content: "Product:"},
			{row: 6, col: 1}: {content: "Max:"},
			{row: 7, col: 1}: {content: "Min:"},

			{row: 1, col: 2}: {content: "12345"},
			{row: 2, col: 2}: {content: "67890"},

			{row: 4, col: 2}: {content: "=SUM(B1:B2)"},
			{row: 5, col: 2}: {content: "=PRODUCT(B1:B2)"},
			{row: 6, col: 2}: {content: "=MAX(B1:B2)"},
			{row: 7, col: 2}: {content: "=MIN(B1:B2)"},
		},
		cursor: Position{row: 1, col: 1},
		selection: []Position{},
	}
}

package main

import "github.com/charmbracelet/lipgloss"

var ThDeselected = lipgloss.NewStyle().
	Foreground(darkGrey).
	Background(hilite).
	Align(lipgloss.Center).
	PaddingLeft(1).
	PaddingRight(1)

var ThSelected = lipgloss.NewStyle().
	Foreground(hilite).
	Background(darkGrey).
	Align(lipgloss.Center).
	PaddingLeft(1).
	PaddingRight(1)

var TrDeselected = lipgloss.NewStyle().
	Foreground(darkGrey).
	Background(hilite).
	Align(lipgloss.Center).
	PaddingLeft(1).
	PaddingRight(1).
	Width(firstColWidth)

var TrSelected = lipgloss.NewStyle().
	Foreground(hilite).
	Background(darkGrey).
	Align(lipgloss.Center).
	PaddingLeft(1).
	PaddingRight(1).
	Width(firstColWidth)

var CursorSelected = lipgloss.NewStyle().
	Foreground(darkGrey).
	Background(hilite).
	PaddingLeft(1).
	PaddingRight(1)

var CursorEditMode = lipgloss.NewStyle().
	Foreground(darkGrey).
	Background(white).
	PaddingLeft(1).
	PaddingRight(1)

var CursorDeselected = lipgloss.NewStyle().
	PaddingLeft(1).
	PaddingRight(1)

var CellReferenced = lipgloss.NewStyle().
	Foreground(white).
	Background(darkGrey).
	PaddingLeft(1).
	PaddingRight(1)

// var Selected = lipgloss.NewStyle().
// 	Background(black).
// 	PaddingLeft(1).
// 	PaddingRight(1)

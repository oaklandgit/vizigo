package main

import "github.com/charmbracelet/lipgloss"

var ThDeselected = lipgloss.NewStyle().
	Foreground(black).
	Background(hilite).
	Align(lipgloss.Center).
	PaddingLeft(1).
	PaddingRight(1)

var ThSelected = lipgloss.NewStyle().
	Foreground(hilite).
	Background(black).
	Align(lipgloss.Center).
	PaddingLeft(1).
	PaddingRight(1)

var TrDeselected = lipgloss.NewStyle().
	Foreground(black).
	Background(hilite).
	Align(lipgloss.Center).
	PaddingLeft(1).
	PaddingRight(1).
	Width(firstColWidth)

var TrSelected = lipgloss.NewStyle().
	Foreground(hilite).
	Background(black).
	Align(lipgloss.Center).
	PaddingLeft(1).
	PaddingRight(1).
	Width(firstColWidth)

var CursorSelected = lipgloss.NewStyle().
	Foreground(black).
	Background(hilite).
	PaddingLeft(1).
	PaddingRight(1)

var CursorEditMode = lipgloss.NewStyle().
	Foreground(black).
	Background(white).
	PaddingLeft(1).
	PaddingRight(1)

var CursorDeselected = lipgloss.NewStyle().
	PaddingLeft(1).
	PaddingRight(1)

var Selected = lipgloss.NewStyle().
	Background(black).
	PaddingLeft(1).
	PaddingRight(1)

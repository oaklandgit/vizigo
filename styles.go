package main

import "github.com/charmbracelet/lipgloss"

var ThDeselected = lipgloss.NewStyle().
	Foreground(Black).
	Background(Hilite).
	Width(ColWidth).
	Align(lipgloss.Center)

var ThSelected = lipgloss.NewStyle().
	Foreground(Hilite).
	Background(Black).
	Width(ColWidth).
	Align(lipgloss.Center)

var TrDeselected = lipgloss.NewStyle().
	Foreground(Black).
	Background(Hilite).
	Width(FirstColWidth).
	Align(lipgloss.Center)

var TrSelected = lipgloss.NewStyle().
	Foreground(Hilite).
	Background(Black).
	Width(FirstColWidth).
	Align(lipgloss.Center)

var CursorSelected = lipgloss.NewStyle().
	Foreground(Black).
	Background(Hilite).
	Width(ColWidth).
	PaddingLeft(1)

var CursorEditMode = lipgloss.NewStyle().
	Foreground(Black).
	Background(White).
	Width(ColWidth).
	PaddingLeft(1)

var CursorDeselected = lipgloss.NewStyle().
	Width(ColWidth).
	PaddingLeft(1)

var Selected = lipgloss.NewStyle().
	Background(Black).
	Width(ColWidth).
	PaddingLeft(1)

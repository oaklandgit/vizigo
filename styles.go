package main

import "github.com/charmbracelet/lipgloss"

var ThDeselected = lipgloss.NewStyle().
	Foreground(Black).
	Background(Hilite).
	Align(lipgloss.Center).
	PaddingLeft(1).
	PaddingRight(1)

var ThSelected = lipgloss.NewStyle().
	Foreground(Hilite).
	Background(Black).
	Align(lipgloss.Center).
	PaddingLeft(1).
	PaddingRight(1)

var TrDeselected = lipgloss.NewStyle().
	Foreground(Black).
	Background(Hilite).
	Align(lipgloss.Center).
	PaddingLeft(1).
	PaddingRight(1).
	Width(FirstColWidth)

var TrSelected = lipgloss.NewStyle().
	Foreground(Hilite).
	Background(Black).
	Align(lipgloss.Center).
	PaddingLeft(1).
	PaddingRight(1).
	Width(FirstColWidth)

var CursorSelected = lipgloss.NewStyle().
	Foreground(Black).
	Background(Hilite).
	PaddingLeft(1).
	PaddingRight(1)

var CursorEditMode = lipgloss.NewStyle().
	Foreground(Black).
	Background(White).
	PaddingLeft(1).
	PaddingRight(1)

var CursorDeselected = lipgloss.NewStyle().
	PaddingLeft(1).
	PaddingRight(1)

var Selected = lipgloss.NewStyle().
	Background(Black).
	PaddingLeft(1).
	PaddingRight(1)

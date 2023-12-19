package main

import "github.com/charmbracelet/lipgloss"

var ThDeselected = lipgloss.NewStyle().
    Foreground(lipgloss.Color(Gray)).
    Background(lipgloss.Color(Orange)).
	Width(ColWidth).
	Align(lipgloss.Center)

var ThSelected = lipgloss.NewStyle().
    Foreground(lipgloss.Color(Orange)).
    Background(lipgloss.Color(Gray)).
	Width(ColWidth).
	Align(lipgloss.Center)

var TrDeselected = lipgloss.NewStyle().
    Foreground(lipgloss.Color(Gray)).
    Background(lipgloss.Color(Orange)).
	Width(FirstColWidth).
	Align(lipgloss.Center)

var TrSelected = lipgloss.NewStyle().
    Foreground(lipgloss.Color(Orange)).
    Background(lipgloss.Color(Gray)).
	Width(FirstColWidth).
	Align(lipgloss.Center)

var CursorSelected = lipgloss.NewStyle().
    Foreground(lipgloss.Color("#000000")).
    Background(lipgloss.Color(Orange)).
	Width(ColWidth).
	PaddingLeft(1)

var CursorDeselected = lipgloss.NewStyle().
	Width(ColWidth).
	PaddingLeft(1)

var Blue = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#0000FF"))

var Green = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#00FF00"))

var Red = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#FF0000"))

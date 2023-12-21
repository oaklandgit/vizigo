package main

import "github.com/charmbracelet/lipgloss"

const (
	Rows      		= 12
	Cols      		= 6
	HOffset			= 1
	VOffset 		= 1
	ColWidth 		= 12
	FirstColWidth 	= 3
	Hilite 			= lipgloss.Color("72")
	Gray 			= lipgloss.Color("243")
	Black 			= lipgloss.Color("0")
	White 			= lipgloss.Color("15")
	
	HelpText = "move: [→ ← ↑ ↓], edit: [enter], copy: [⌃c], paste: [⌃v], save: [⌃s], exit: [⌃x]"
)
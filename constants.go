package main

import "github.com/charmbracelet/lipgloss"

const (
	defaultFilename 		= "untitled"
	fileExtension	 		= ".viz"
	defaultCols				= 6
	defaultRows				= 12
	hOffset        			= 1
	vOffset        			= 1
	minColWidth    			= 8
	maxEntryLength 			= 22
	firstColWidth  			= 4
	hilite         			= lipgloss.Color("72")
	gray           			= lipgloss.Color("243")
	black          			= lipgloss.Color("0")
	white          			= lipgloss.Color("15")
	globalPrecisionLimit 	= 5 // (decimal places)
)

var helpText = map[string]string{
	"move":  "→ ← ↑ ↓",
	"edit":  "enter, tab",
	"copy":  "ctrl-c",
	"paste": "ctrl-v",
	"save":  "ctrl-s",
	"quit":  "ctrl-q",
}

var helpTextKeys = []string{"move", "edit", "copy", "paste", "save", "quit"}

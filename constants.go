package main

import "github.com/charmbracelet/lipgloss"

const (
	defaultFilename 		= "untitled"
	fileExtension	 		= ".viz"
	defaultCols				= 20
	defaultRows				= 20
	viewportCols			= 8
	viewportRows			= 12
	minColWidth    			= 8
	maxEntryLength 			= 22
	firstColWidth  			= 4
	white          			= lipgloss.Color("15")
	darkGrey          		= lipgloss.Color("0")
	hilite         			= lipgloss.Color("72")
	globalPrecisionLimit 	= 5 // (decimal places)
)

var helpText = `

==== HELP ====

values              : enter, then type a value
labels              : enter, then type a label
formulae            : e.g. =SUM(A1:B2, B3, C5)

move                : → ← ↑ ↓
edit                : enter, tab
copy                : ctrl-c
copy (value only)   : ctrl-n
paste               : ctrl-v
save                : ctrl-s
quit                : ctrl-q

`
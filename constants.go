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

var showHelp = false

var helpText = `
flags (on start)      -c <cols> -r <rows>
                      -vc <viewport cols> -vr <viewport rows>
                      -f <filename>

enter values          <enter> or <tab> then type numbers
enter labels          <enter> or <tab> then type letters
enter expression      <enter> or <tab> then type =expression
                      e.g. =min(A3:B5)
                      use any EXPR expression (see expr-lang.org)

move                  → ← ↑ ↓
edit                  <enter> or <tab>
                      <esc> to exit edit mode
copy                  ⌃c
copy value only       ⌃n
paste                 ⌃v

save                  ⌃s
quit                  ⌃q
`
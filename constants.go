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
enter values          <enter> or <tab> then type numbers
enter labels          <enter> or <tab> then type letters
enter formulae        <enter> or <tab> then type =FORMULA()
                      example: =SUM(A1:B2, B3, C5)
                      also try: =PROD(), =AVG(), =MIN(), =MAX(), =COUNT()

move                  → ← ↑ ↓
edit                  <enter> or <tab>
                      <esc> to exit edit mode
copy                  ⌃c
copy value only       ⌃n
paste                 ⌃v

save                  ⌃s
quit                  ⌃q

flags (on start)      -c <cols> -r <rows> -vc <viewport cols> -vr <viewport rows> -f <filename>
`
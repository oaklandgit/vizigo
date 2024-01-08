# vizigo

**WORK IN PROGRESS**

A command-line spreadsheet written in Go and [BubbleTea](https://github.com/charmbracelet/bubbletea)

![demo](demo.gif)

### TO DO

- select ranges visually
- implement simple arithmetic (not just functions)
- autosuggest formulae
- export as csv
- complete unit tests

### DONE

- viewport and scrolling
- implement SUM, PRODUCT, AVERAGE, MIN, MAX, COUNT
- comma-separated ranges
- hilite references to selected cell
- undo / redo
- load file
- some unit tests
- round results to the most-precise of its operands
- filepath/filename for save
- flags -r (rows) -c (columns) -f (filename)

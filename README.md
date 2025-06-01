# vizigo

**WORK IN PROGRESS**

An intuitive TUI (Terminal User Interface) spreadsheet written in Go and [BubbleTea](https://github.com/charmbracelet/bubbletea) that leverages the powerful expression language [Expr](https://expr-lang.org/).

![demo](demo.gif)

### How to use

```
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

```

### TO DO

- implement functions not available in EXPR (such as prod, sum, avg)
- select ranges visually
- autosuggest formulae
- export as csv
- complete unit tests

### DONE

- implement simple arithmetic (not just functions) // via expr language
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

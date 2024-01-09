# vizigo

**WORK IN PROGRESS**

A command-line spreadsheet written in Go and [BubbleTea](https://github.com/charmbracelet/bubbletea)

![demo](demo.gif)

### How to use

```
flags (on start)      -c <cols> -r <rows>
                      -vc <viewport cols> -vr <viewport rows>
                      -f <filename>

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
```

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

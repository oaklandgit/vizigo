# vizigo

> ⚠️ **WORK IN PROGRESS** – This project is incomplete and not ready for general use.

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
                      (examples below)

move                  → ← ↑ ↓
edit                  <enter> or <tab>
                      <esc> to exit edit mode
copy                  ⌃c
copy value only       ⌃n
paste                 ⌃v

save                  ⌃s
quit                  ⌃q

```

### Math Functions

```
sum, product, max, min, average, count, abs, ceil, floor, round, int, float
```

### String Functions

```
upper, lower, trim, trimPrefix, trimSuffix, replace, split, join, len, string, contains, indexOf
```

### Operators

```
Arithmetic: + - \* / % \*\* (power)
Comparison: == != < <= > >=
Logic: && || !`
Ternary: condition ? a : b
String concat: "hello" + " world"
```

### Examples

```
=A1+B1                        add two cells
=A1*1.08                      multiply (e.g. add tax)
=SUM(B1:B10)                  sum a range
=AVERAGE(B1:B10)              average a range
=MAX(B1:B10)-MIN(B1:B10)      spread of a range
=COUNT(B1:B10)                count cells in a range
=ABS(A1-B1)                   absolute difference
=A1 > 100 ? "over" : "under"  conditional label
=UPPER(A1)                    uppercase a label
=LOWER(A1)+" "+LOWER(B1)      concatenate labels
=LEN(A1)                      length of a label
```

### TO DO

- open example file with a flag (e.g. vizigo --example or -e)
- improve help display. Modal window?
- implement save as: Filename: [ mysheet    ].viz  Path: [ ~/.vizigo/sheets/ ]
- select ranges visually
- autosuggest formulae
- export as csv
- complete unit tests

### DONE

- fix =SUM(1,2,3) resulting in error due to float64 casting
- implement functions not available in EXPR (such as prod, sum, avg)
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

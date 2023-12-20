package main

import "unicode"

func ColumnToLetters(n int) string {
	var result string
	for n > 0 {
		remainder := (n - 1) % 26
		result = string('A' + remainder) + result
		n = (n - 1) / 26
	}
	return result
}

func LettersToColumn(s string) int {
	var result int
	for _, c := range s {
		result *= 26
		result += int(c) - 'A' + 1
	}
	return result
}

func GetCellContent(g Grid, p Position) string {
	return g.cells[p].content
}

func SetCellContent(g *Grid, p Position, content string) {
	g.cells[p] = Cell{content: content}
}

func SplitAlphaNumeric(s string) (alphaPart string, numericPart string) {
    splitIndex := -1
    for i, char := range s {
        if unicode.IsDigit(char) {
            splitIndex = i
            break
        }
    }
    if splitIndex != -1 {
        return s[:splitIndex], s[splitIndex:]
    }
    return s, ""
}
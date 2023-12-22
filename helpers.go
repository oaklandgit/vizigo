package main

import (
	"fmt"
	"unicode"
)

func UnderlineChar(s string, i int) string {

	if len(s) == 0 {
		return ""
	}

	start := s[:i]
    underline := s[i : i+1]
    end := s[i+1:]

	return fmt.Sprintf("%s\033[4m%s\033[0m%s", start, underline, end)
}

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
package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

func padStringToCenter(s string, width int) string {
	if len(s) >= width {
		return s
	}
	leftPadding := (width - len(s)) / 2
	rightPadding := width - len(s) - leftPadding
	return strings.Repeat(" ", leftPadding) + s + strings.Repeat(" ", rightPadding)
}

func splitStringAt(s string, i int) (string, string, string, error) {
	if i < 0 || i > len(s) {
		return "", "", "", fmt.Errorf("can't split string at %d", i)
	}
	return s[:i], s[i : i+1], s[i+1:], nil
}

func underlineChar(s string, i int) string {

	if i < 0 {
		return s
	}

	start, underline, end, error := splitStringAt(s, i)
	if error != nil {
		log.Fatal(error)
	}

	return fmt.Sprintf("%s\033[4m%s\033[0m%s", start, underline, end)
	
}

func columnToLetters(n int) string {
	var result string
	for n > 0 {
		remainder := (n - 1) % 26
		result = string(rune('A'+remainder)) + result
		n = (n - 1) / 26
	}
	return result
}

func lettersToColumn(s string) int {
	var result int
	for _, c := range s {
		result *= 26
		result += int(c) - 'A' + 1
	}
	return result
}

func splitAlphaNumeric(s string) (alphaPart string, numericPart string) {
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

func alphaNumericToPosition(s string) Position {
	alphaPart, numericPart := splitAlphaNumeric(s)
	col := lettersToColumn(alphaPart)
	row, _ := strconv.Atoi(numericPart)
	return Position{row: row, col: col}
}

func maxPrecision(operands []float64) int {
	
	max := 0

	for _, operand := range operands {
		
		str := strconv.FormatFloat(operand, 'f', -1, 64)
		parts := strings.Split(str, ".")

		if len(parts) == 2 {
			decPlaces := len(parts[1])
			if decPlaces > max {
				max = decPlaces
			}
		}

	}
	return max
}
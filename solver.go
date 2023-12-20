package main

import (
	"math"
	"regexp"
	"strconv"
)


func (g *Grid) Compute(s string) string {

	pattern := `=([A-Z]+)\(([A-Z]+\d+):([A-Z]+\d+)\)`
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(s)

	if matches == nil {
		return s
	}


	startCell := matches[2]
	endCell := matches[3]

	operand1, _ := strconv.ParseFloat(g.CellFromString(startCell).content, 64)
	operand2, _ := strconv.ParseFloat(g.CellFromString(endCell).content, 64)

	result := 0.0

	switch matches[1] {
	case "SUM":
		result = operand1 + operand2
	case "PROD":
		result = operand1 * operand2
	case "DIFF":
		result = operand1 - operand2
	case "MAX":
		result = math.Max(operand1, operand2)
	case "MIN":
		result = math.Min(operand1, operand2)
	}

	return strconv.FormatFloat(result, 'f', -1, 64)

}
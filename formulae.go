package main

import (
	"strconv"
	"strings"
	"regexp"
)

// Arguments may contain quoted strings (text cells), which can hold parens or commas.
var customFuncRe = regexp.MustCompile(`\b(sum|product|max|min|average|count)\(((?:[^()"']|"(?:[^"\\]|\\.)*"|'(?:[^'\\]|\\.)*')*)\)`)

// blankCell stands in for a referenced empty cell. Aggregate functions skip it,
// as spreadsheets do; anywhere else it becomes 0.
const blankCell = "_blank_"

// splitArgs splits a comma-separated argument list, ignoring commas inside quotes.
func splitArgs(argsStr string) []string {
	var args []string
	start := 0
	var quote byte
	for i := 0; i < len(argsStr); i++ {
		c := argsStr[i]
		switch {
		case quote != 0 && c == '\\':
			i++
		case quote != 0 && c == quote:
			quote = 0
		case quote == 0 && (c == '"' || c == '\''):
			quote = c
		case quote == 0 && c == ',':
			args = append(args, argsStr[start:i])
			start = i + 1
		}
	}
	return append(args, argsStr[start:])
}

func formatNumber(f float64) string {
	return strconv.FormatFloat(f, 'g', 10, 64)
}

// evaluateCustomFunctions replaces calls to our aggregation functions with
// their computed values. Applied repeatedly so nested calls resolve inside-out.
func evaluateCustomFunctions(expression string) string {
	for {
		next := customFuncRe.ReplaceAllStringFunc(expression, func(m string) string {
			parts := customFuncRe.FindStringSubmatch(m)
			if len(parts) < 3 {
				return m
			}
			name, argsStr := parts[1], parts[2]

			// Like spreadsheets, only numbers count: blanks, text and booleans
			// are skipped, but an error anywhere in the range propagates.
			var args []float64
			for _, a := range splitArgs(argsStr) {
				a = strings.TrimSpace(a)
				switch {
				case a == "" || a == blankCell || a == "true" || a == "false":
					continue
				case a[0] == '"' || a[0] == '\'':
					continue
				}
				f, err := strconv.ParseFloat(a, 64)
				if err != nil {
					return errorText
				}
				args = append(args, f)
			}
			if len(args) == 0 {
				if name == "average" {
					return errorText // no numbers to average (#DIV/0! in other spreadsheets)
				}
				return "0.0"
			}

			var result float64
			switch name {
			case "sum":
				result = sum(args...)
			case "product":
				result = product(args...)
			case "max":
				result = max(args...)
			case "min":
				result = min(args...)
			case "average":
				result = average(args...)
			case "count":
				result = count(args...)
			}
			return formatNumber(result)
		})

		if next == expression {
			break
		}
		expression = next
	}
	return expression
}

func sum(operands ...float64) float64 {
	total := 0.0
	for _, f := range operands {
		total += f
	}
	return total
}

func product(operands ...float64) float64 {
	total := 1.0
	for _, f := range operands {
		total *= f
	}
	return total
}

func max(operands ...float64) float64 {
	m := operands[0]
	for _, f := range operands[1:] {
		if f > m {
			m = f
		}
	}
	return m
}

func min(operands ...float64) float64 {
	m := operands[0]
	for _, f := range operands[1:] {
		if f < m {
			m = f
		}
	}
	return m
}

func average(operands ...float64) float64 {
	return sum(operands...) / float64(len(operands))
}

func count(operands ...float64) float64 {
	return float64(len(operands))
}

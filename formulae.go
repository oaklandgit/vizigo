package main

import (
	"strconv"
	"strings"
	"regexp"
)

var customFuncRe = regexp.MustCompile(`\b(sum|product|max|min|average|count)\(([^()]*)\)`)

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

			var args []float64
			for _, a := range strings.Split(argsStr, ",") {
				f, err := strconv.ParseFloat(strings.TrimSpace(a), 64)
				if err != nil {
					return errorText
				}
				args = append(args, f)
			}
			if len(args) == 0 {
				return errorText
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

package main

import "strconv"

func Solver(s string) string {

	// formula
	if s[0] == '=' {
		return Red.Render(s)
	}

	// number
	if _, err := strconv.Atoi(s); err == nil {
		return Green.Render(s)
	}

	// string
	return Blue.Render(s)
	
}
package main

type Position struct {
	row, col int
}

type Cell struct {
	content  string
}

type Grid struct {
	cells map[Position]Cell
	cursor Position
}


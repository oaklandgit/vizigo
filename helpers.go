package main

func IntToLetters(n int) string {
	var result string
	for n > 0 {
		remainder := (n - 1) % 26
		result = string('A' + remainder) + result
		n = (n - 1) / 26
	}
	return result
}

func GetCellContent(g Grid, p Position) string {
	return g.cells[p].content
}

func SetCellContent(g *Grid, p Position, content string) {
	g.cells[p] = Cell{content: content}
}
package main

type cursor struct {
	vector
	editMode  bool
	editIndex int
	clipboard string
}

func (c *cursor) copy(g *grid) {
	c.editMode = false
	c.clipboard = c.getCellContent(g, false)
}

func (c *cursor) copyValue(g *grid) {
	c.editMode = false
	c.clipboard = c.getCellContent(g, true)
}

func (c *cursor) paste(g *grid) {
	c.editMode = false
	c.setCellContent(g, c.clipboard)
	g.saveForUndo()
}

func (c *cursor) enter(g *grid) {
	if c.editMode {
		c.editMode = false
		c.editIndex = -1
		if c.row < g.size.row {
			c.row++
		}
		g.saveForUndo()
	} else {
		c.editMode = true
	}
}

func (c *cursor) tab(g *grid) {
	if c.editMode {
		c.editMode = false
		c.editIndex = -1
		if c.col < g.size.col {
			c.col++
		}
		g.saveForUndo()
	} else {
		c.editMode = true
	}
}

func (c *cursor) up(g *grid) {
	c.editMode = false
	c.editIndex = -1
	if c.row > 1 {
		c.row--
		if c.row < g.viewport.offset.row  {
			g.viewport.offset.row--
		}
	}
}

func (c *cursor) down(g *grid) {
	c.editMode = false
	c.editIndex = -1
	if c.row < g.size.row {
		c.row++
		if c.row == g.viewport.size.row + g.viewport.offset.row  {
			g.viewport.offset.row++
		}
	}


}

func (c *cursor) left(g *grid) {
	if !c.editMode && c.col > 1 {
		c.col--

		if c.col < g.viewport.offset.col {
			g.viewport.offset.col--
		}
	} else if c.editMode && c.editIndex > -1 {
			c.editIndex--
	}
}

func (c *cursor) right(g *grid) {
	if !c.editMode && c.col < g.size.col {
		
		c.col++

		if c.col == g.viewport.size.col + g.viewport.offset.col  {
			g.viewport.offset.col++
		}
	} else if c.editMode && c.editIndex < len(c.getCellContent(g, false)) -1 {
		c.editIndex++
	}
}



func (c *cursor) textEntry(g *grid, s string) {
	if !c.editMode || c.editIndex == maxEntryLength {
		return
	}

	c.editIndex++

	before := c.getCellContent(g, false)[:c.editIndex]
	after := c.getCellContent(g, false)[c.editIndex:]

	c.setCellContent(g, before + s + after)

}

func (c *cursor) clear(g *grid) {
	delete(g.cells, c.vector)
	delete(g.computed, c.vector)
	g.saveForUndo()
}

func (c *cursor) backspace(g *grid) {
	if !c.editMode {
		c.clear(g)
		return
	}

	if c.editIndex > -1 {

		before := c.getCellContent(g, false)[:c.editIndex + 1]
		after := c.getCellContent(g, false)[c.editIndex + 1:]
		c.setCellContent(g, before[:len(before) -1] + after)
		c.editIndex--
		
	}
}

func (c *cursor) escape() {
	c.editMode = false
	c.editIndex = -1
}

package main

type cursor struct {
	vector
	editMode  bool
	editIndex int
	clipboard string
}

func (c *cursor) copy(g *grid) {
	c.editMode = false
	c.clipboard = c.vector.getCellContent(g, false)
}

func (c *cursor) copyValue(g *grid) {
	c.editMode = false
	c.clipboard = c.vector.getCellContent(g, true)
}

func (c *cursor) paste(g *grid) {
	c.editMode = false
	c.vector.setCellContent(g, c.clipboard)
	g.saveForUndo()
}

func (c *cursor) enter(g *grid) {
	if c.editMode {
		c.editMode = false
		c.editIndex = -1
		if c.vector.row < g.size.row {
			c.vector.row++
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
		if c.vector.col < g.size.col {
			c.vector.col++
		}
		g.saveForUndo()
	} else {
		c.editMode = true
	}
}

func (c *cursor) up() {
	c.editMode = false
	c.editIndex = -1
	if c.vector.row > 1 {
		c.vector.row--
	}
}

func (c *cursor) down(g *grid) {
	c.editMode = false
	c.editIndex = -1
	if c.vector.row < g.size.row {
		c.vector.row++
	}
}

func (c *cursor) left(g *grid) {
	if !c.editMode && c.vector.col > 1 {
		c.vector.col--

		if c.vector.col < g.viewport.offset.col {
			g.viewport.offset.col--
		}
	} else if c.editMode && c.editIndex > -1 {
			c.editIndex--
	}
}

func (c *cursor) right(g *grid) {
	if !c.editMode && c.vector.col < g.viewport.size.col {
		
		c.vector.col++

		if c.vector.col >= g.viewport.size.col  {
			g.viewport.offset.col++
		}
	} else if c.editMode && c.editIndex < len(c.vector.getCellContent(g, false)) -1 {
		c.editIndex++
	}
}



func (c *cursor) textEntry(g *grid, s string) {
	if !c.editMode || c.editIndex == maxEntryLength {
		return
	}

	c.editIndex++

	before := c.vector.getCellContent(g, false)[:c.editIndex]
	after := c.vector.getCellContent(g, false)[c.editIndex:]

	c.vector.setCellContent(g, before + s + after)

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

		before := c.vector.getCellContent(g, false)[:c.editIndex + 1]
		after := c.vector.getCellContent(g, false)[c.editIndex + 1:]
		c.vector.setCellContent(g, before[:len(before) -1] + after)
		c.editIndex--
		
	}
}

func (c *cursor) escape() {
	c.editMode = false
	c.editIndex = -1
}

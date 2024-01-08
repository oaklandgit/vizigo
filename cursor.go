package main

type cursor struct {
	vector
	editMode  bool
	editIndex int
	clipboard string
}

func (c *cursor) Copy(g *grid) {
	c.editMode = false
	c.clipboard = c.vector.GetcellContent(g, false)
}

func (c *cursor) CopyValue(g *grid) {
	c.editMode = false
	c.clipboard = c.vector.GetcellContent(g, true)
}

func (c *cursor) Paste(g *grid) {
	c.editMode = false
	c.vector.SetcellContent(g, c.clipboard)
	g.saveForUndo()
}

func (c *cursor) Enter(g *grid) {
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

func (c *cursor) Tab(g *grid) {
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

func (c *cursor) Up() {
	c.editMode = false
	c.editIndex = -1
	if c.vector.row > 1 {
		c.vector.row--
	}
}

func (c *cursor) Down(g *grid) {
	c.editMode = false
	c.editIndex = -1
	if c.vector.row < g.size.row {
		c.vector.row++
	}
}

func (c *cursor) Left(g *grid) {
	if !c.editMode && c.vector.col > 1 {
		c.vector.col--

		if c.vector.col < g.viewport.offset.col {
			g.viewport.offset.col--
		}
	} else if c.editMode && c.editIndex > -1 {
			c.editIndex--
	}
}

func (c *cursor) Right(g *grid) {
	if !c.editMode && c.vector.col < g.viewport.size.col {
		
		c.vector.col++

		if c.vector.col >= g.viewport.size.col  {
			g.viewport.offset.col++
		}
	} else if c.editMode && c.editIndex < len(c.vector.GetcellContent(g, false)) -1 {
		c.editIndex++
	}
}



func (c *cursor) TextEntry(g *grid, s string) {
	if !c.editMode || c.editIndex == maxEntryLength {
		return
	}

	c.editIndex++

	before := c.vector.GetcellContent(g, false)[:c.editIndex]
	after := c.vector.GetcellContent(g, false)[c.editIndex:]

	c.vector.SetcellContent(g, before + s + after)

}

func (c *cursor) Clear(g *grid) {
	delete(g.cells, c.vector)
	delete(g.computed, c.vector)
	g.saveForUndo()
}

func (c *cursor) Backspace(g *grid) {
	if !c.editMode {
		c.Clear(g)
		return
	}

	if c.editIndex > -1 {

		before := c.vector.GetcellContent(g, false)[:c.editIndex + 1]
		after := c.vector.GetcellContent(g, false)[c.editIndex + 1:]
		c.vector.SetcellContent(g, before[:len(before) -1] + after)
		c.editIndex--
		
	}
}

func (c *cursor) Escape() {
	c.editMode = false
	c.editIndex = -1
}

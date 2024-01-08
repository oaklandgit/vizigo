package main

type Cursor struct {
	Vector
	editMode  bool
	editIndex int
	clipboard string
}

func (c *Cursor) Copy(g *Grid) {
	c.editMode = false
	c.clipboard = c.Vector.GetCellContent(g, false)
}

func (c *Cursor) CopyValue(g *Grid) {
	c.editMode = false
	c.clipboard = c.Vector.GetCellContent(g, true)
}

func (c *Cursor) Paste(g *Grid) {
	c.editMode = false
	c.Vector.SetCellContent(g, c.clipboard)
	g.SaveForUndo()
}

func (c *Cursor) Enter(g *Grid) {
	if c.editMode {
		c.editMode = false
		c.editIndex = -1
		if c.Vector.row < g.size.row {
			c.Vector.row++
		}
		g.SaveForUndo()
	} else {
		c.editMode = true
	}
}

func (c *Cursor) Tab(g *Grid) {
	if c.editMode {
		c.editMode = false
		c.editIndex = -1
		if c.Vector.col < g.size.col {
			c.Vector.col++
		}
		g.SaveForUndo()
	} else {
		c.editMode = true
	}
}

func (c *Cursor) Up() {
	c.editMode = false
	c.editIndex = -1
	if c.Vector.row > 1 {
		c.Vector.row--
	}
}

func (c *Cursor) Down(g *Grid) {
	c.editMode = false
	c.editIndex = -1
	if c.Vector.row < g.size.row {
		c.Vector.row++
	}
}

func (c *Cursor) Left(g *Grid) {
	if !c.editMode && c.Vector.col > 1 {
		c.Vector.col--

		if c.Vector.col < g.viewport.offset.col {
			g.viewport.offset.col--
		}
	} else if c.editMode && c.editIndex > -1 {
			c.editIndex--
	}
}

func (c *Cursor) Right(g *Grid) {
	if !c.editMode && c.Vector.col < g.viewport.size.col {
		
		c.Vector.col++

		if c.Vector.col >= g.viewport.size.col  {
			g.viewport.offset.col++
		}
	} else if c.editMode && c.editIndex < len(c.Vector.GetCellContent(g, false)) -1 {
		c.editIndex++
	}
}



func (c *Cursor) TextEntry(g *Grid, s string) {
	if !c.editMode || c.editIndex == maxEntryLength {
		return
	}

	c.editIndex++

	before := c.Vector.GetCellContent(g, false)[:c.editIndex]
	after := c.Vector.GetCellContent(g, false)[c.editIndex:]

	c.Vector.SetCellContent(g, before + s + after)

}

func (c *Cursor) Clear(g *Grid) {
	// c.Vector.SetCellContent(g, "")
	delete(g.cells, c.Vector)
	delete(g.computed, c.Vector)
	g.SaveForUndo()
}

func (c *Cursor) Backspace(g *Grid) {
	if !c.editMode {
		c.Clear(g)
		return
	}

	if c.editIndex > -1 {

		before := c.Vector.GetCellContent(g, false)[:c.editIndex + 1]
		after := c.Vector.GetCellContent(g, false)[c.editIndex + 1:]
		c.Vector.SetCellContent(g, before[:len(before) -1] + after)
		c.editIndex--
		
	}
}

func (c *Cursor) Escape() {
	c.editMode = false
	c.editIndex = -1
}

package main

type Cursor struct {
	position  Position
	editMode  bool
	editIndex int
	clipboard string
}

func (c *Cursor) Copy(g *Grid) {
	c.editMode = false
	c.clipboard = c.position.GetCellContent(g)
}

func (c *Cursor) Paste(g *Grid) {
	c.editMode = false
	c.position.SetCellContent(g, c.clipboard)
	g.SaveForUndo()
}

func (c *Cursor) ToggleEditMode(g *Grid) {
	if c.editMode {
		c.editMode = false
		c.editIndex = -1
		if c.position.row < g.size.row {
			c.position.row++
		}
		g.SaveForUndo()
	} else {
		c.editMode = true
	}
}

func (c *Cursor) Up() {
	c.editMode = false
	c.editIndex = -1
	if c.position.row > 1 {
		c.position.row--
	}
}

func (c *Cursor) Down(g *Grid) {
	c.editMode = false
	c.editIndex = -1
	if c.position.row < g.size.row {
		c.position.row++
	}
}

func (c *Cursor) Left() {
	if !c.editMode && c.position.col > 1 {
		c.position.col--
	} else if c.editMode && c.editIndex > -1 {
			c.editIndex--
	}
}

func (c *Cursor) Right(g *Grid) {
	if !c.editMode && c.position.col < g.size.col {
		c.position.col++
	} else if c.editMode && c.editIndex < len(c.position.GetCellContent(g)) -1 {
		c.editIndex++
	}
}

func (c *Cursor) Entry(g *Grid, s string) {
	if !c.editMode || c.editIndex == maxEntryLength {
		return
	}

	c.editIndex++

	before := c.position.GetCellContent(g)[:c.editIndex]
	after := c.position.GetCellContent(g)[c.editIndex:]

	c.position.SetCellContent(g, before + s + after)

}

func (c *Cursor) Clear(g *Grid) {
	c.position.SetCellContent(g, "")
	g.SaveForUndo()
}

func (c *Cursor) Backspace(g *Grid) {
	if !c.editMode {
		c.Clear(g)
		return
	}

	if c.editIndex > -1 {

		before := c.position.GetCellContent(g)[:c.editIndex + 1]
		after := c.position.GetCellContent(g)[c.editIndex + 1:]
		c.position.SetCellContent(g, before[:len(before) -1] + after)
		c.editIndex--
		
	}
}

func (c *Cursor) Escape() {
	c.editMode = false
	c.editIndex = -1
}

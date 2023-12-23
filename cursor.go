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
}

func (c *Cursor) ToggleEditMode() {
	if c.editMode {
		c.editMode = false
		c.editIndex = 0
		if c.position.row < Rows-1 {
			c.position.row++
		}
	} else {
		c.editMode = true
	}
}

func (c *Cursor) Up() {
	c.editMode = false
	c.editIndex = 0
	if c.position.row > 1 {
		c.position.row--
	}
}

func (c *Cursor) Down() {
	c.editMode = false
	c.editIndex = 0
	if c.position.row < Rows-1 {
		c.position.row++
	}
}

func (c *Cursor) Left() {
	if !c.editMode && c.position.col > 1 {
		c.position.col--
	} else if c.editMode && c.editIndex > 0 {
		c.editIndex--
	}
}

func (c *Cursor) Right() {
	if !c.editMode && c.position.col < Cols-1 {
		c.position.col++
	} else if c.editMode && c.editIndex < ColWidth-1 {
		c.editIndex++
	}
}

func (c *Cursor) Entry(g *Grid, s string) {
	if !c.editMode || c.editIndex == MaxEntryLength {
		return
	}
	was := c.position.GetCellContent(g)
	c.position.SetCellContent(g, was+s)
	c.editIndex++
}

func (c *Cursor) Backspace(g *Grid) {
	if !c.editMode || c.editIndex == 0 {
		c.position.SetCellContent(g, "")
	} else {
		was := c.position.GetCellContent(g)
		if len(was) > 0 {
			c.position.SetCellContent(g, was[:len(was)-1])
			c.editIndex--
		}
	}
}

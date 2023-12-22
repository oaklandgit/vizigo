package main

type Cursor struct {
	position Position
	editMode bool
	editIndex int
	clipboard string
}

func (c *Cursor) backspace(g *Grid) {
	if c.editMode {
		SetCellContent(g, c.position, "")
	} else {
		was := GetCellContent(g, c.position)
		if len(was) > 0 {
			SetCellContent(g, g.cursor.position, was[:len(was)-1])
			g.cursor.editIndex--
		}
	}
}

func (c *Cursor) Copy(g *Grid) {
	c.editMode = false
	c.clipboard = GetCellContent(g, c.position)
}

func (c *Cursor) Paste(g *Grid) {
	c.editMode = false
	SetCellContent(g, c.position, c.clipboard)
}

func (c *Cursor) ToggleEditMode() {
	if c.editMode {
		c.editMode = false
		c.editIndex = 0
		if c.position.row < Rows - 1 {
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
	if c.position.row < Rows - 1 {
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
	if !c.editMode && c.position.col < Cols - 1 {
		c.position.col++
	} else if c.editMode && c.editIndex < ColWidth - 1 {
		c.editIndex++
	}
}

// if !g.cursor.editMode {
// 	return g, nil
// }
// was := GetCellContent(&g, g.cursor.position)
// SetCellContent(&g, g.cursor.position, was + msg.String())
// g.cursor.editIndex++

func (c *Cursor) Entry(g *Grid, s string) {
	if !c.editMode || c.editIndex == MaxEntryLength {
		return
	}
	was := GetCellContent(g, c.position)
	SetCellContent(g, c.position, was + s)
	c.editIndex++
}

func (c *Cursor) Backspace(g *Grid) {
	if !c.editMode || c.editIndex == 0 {
		SetCellContent(g, c.position, "")
	} else {
		was := GetCellContent(g, c.position)
		if len(was) > 0 {
			SetCellContent(g, c.position, was[:len(was)-1])
			c.editIndex--
		}
	}
}
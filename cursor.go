package main

import "unicode/utf8"

type cursor struct {
	vector
	editMode  bool
	editIndex int
	clipboard string
}

func (c *cursor) copy(s *sheet) {
	c.editMode = false
	c.clipboard = c.getCellContent(s, false)
}

func (c *cursor) copyValue(s *sheet) {
	c.editMode = false
	c.clipboard = c.getCellContent(s, true)
}

func (c *cursor) paste(s *sheet) {
	c.editMode = false
	c.setCellContent(s, c.clipboard)
	s.saveForUndo()
}

func (c *cursor) enter(s *sheet) {
	if c.editMode {
		c.editMode = false
		c.editIndex = -1
		if c.row < s.size.row {
			c.row++
		}
		s.saveForUndo()
	} else {
		c.editMode = true
	}
}

func (c *cursor) tab(s *sheet) {
	if c.editMode {
		c.editMode = false
		c.editIndex = -1
		if c.col < s.size.col {
			c.col++
		}
		s.saveForUndo()
	} else {
		c.editMode = true
	}
}

func (c *cursor) up(s *sheet) {
	c.editMode = false
	c.editIndex = -1
	if c.row > 1 {
		c.row--
		if c.row < s.viewport.offset.row  {
			s.viewport.offset.row--
		}
	}
}

func (c *cursor) down(s *sheet) {
	c.editMode = false
	c.editIndex = -1
	if c.row < s.size.row {
		c.row++
		if c.row == s.viewport.size.row + s.viewport.offset.row  {
			s.viewport.offset.row++
		}
	}


}

func (c *cursor) left(s *sheet) {
	if !c.editMode && c.col > 1 {
		c.col--

		if c.col < s.viewport.offset.col {
			s.viewport.offset.col--
		}
	} else if c.editMode && c.editIndex > -1 {
			c.editIndex--
	}
}

func (c *cursor) right(s *sheet) {
	if !c.editMode && c.col < s.size.col {
		
		c.col++

		if c.col == s.viewport.size.col + s.viewport.offset.col  {
			s.viewport.offset.col++
		}
	} else if c.editMode && c.editIndex < utf8.RuneCountInString(c.getCellContent(s, false)) -1 {
		c.editIndex++
	}
}



func (c *cursor) textEntry(s *sheet, str string) {
	if !c.editMode || c.editIndex == maxEntryLength {
		return
	}

	c.editIndex++

	before := c.getCellContent(s, false)[:c.editIndex]
	after := c.getCellContent(s, false)[c.editIndex:]

	c.setCellContent(s, before + str + after)

}

func (c *cursor) clear(s *sheet) {
	delete(s.cells, c.vector)
	delete(s.computed, c.vector)
	s.saveForUndo()
}

func (c *cursor) backspace(s *sheet) {
	if !c.editMode {
		c.clear(s)
		return
	}

	if c.editIndex > -1 {

		before := c.getCellContent(s, false)[:c.editIndex + 1]
		after := c.getCellContent(s, false)[c.editIndex + 1:]
		c.setCellContent(s, before[:utf8.RuneCountInString(before) -1] + after)
		c.editIndex--
		
	}
}

func (c *cursor) escape() {
	c.editMode = false
	c.editIndex = -1
}

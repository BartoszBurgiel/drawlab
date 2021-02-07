package canvas

// Rect draws a rectange at the given
// coordinates, with given sizes
func (c *Canvas) Rect(x1, y1, width, height int) {

	// get current char
	currChar := c.GetChar()

	// set char to line char
	c.SetChar('|')

	// Lines cannot be drawn since the m = +Inf
	// The points are drawn manually
	for i := 0; i < height; i++ {

		// left line
		c.Point(x1, y1+i)

		// right line
		c.Point(x1+width, y1+i)
	}

	// set char to line char
	c.SetChar('-')

	// top line
	c.Line(x1, y1, x1+width, y1)

	// bottom line
	c.Line(x1, y1+height, x1+width, y1+height)

	// set current char to the initial char
	c.SetChar(currChar)
}

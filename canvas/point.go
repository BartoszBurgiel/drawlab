package canvas

// Point sets the current character to
// this given point
func (c *Canvas) Point(x, y int) {
	c.Fields[c.Height-y-1][x] = c.Char
}

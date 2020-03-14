package canvas

// Point sets the current character to
// this given point
func (c *Canvas) Point(x, y int) {

	// check out of border error
	if x > 0 && y > 0 {
		if x < c.Width-1 && y < c.Height-1 {
			c.Fields[y][x] = c.color + c.Char
		}
	}
}

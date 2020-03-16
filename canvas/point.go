package canvas

// Point sets the current character to
// this given point
func (c *Canvas) Point(x, y int) {

	// check out of border error
	if x > 0 && y > 0 {
		if x < c.Width && y <= c.Height {
			c.Fields[y-1][x] = c.color + c.Char
		}
	}
}

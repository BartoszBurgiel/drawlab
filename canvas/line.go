package canvas

// Line connects two given coordinates on the canvas
func (c *Canvas) Line(x1, y1, x2, y2 int) {

	if f, ok := c.lineFunction(x1, y1, x2, y2); !ok {

		// set all affected fields to c.Char
		for i := x1; i <= x2; i++ {

			c.Point(i, f(i))

		}
	} else {

		// determine start and end
		start := y1
		end := y2
		if start > end {
			end = y1
			start = y2
		}

		for i := start; i < end; i++ {
			c.Point(x1, i)
		}
	}
}

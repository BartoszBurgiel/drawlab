package canvas

// Triangle draws a triangle with the points on the
// provided coordinates
func (c *Canvas) Triangle(x1, y1, x2, y2, x3, y3 int) {
	c.Line(x1, y1, x2, y2)
	c.Line(x2, y2, x3, y3)
	c.Line(x1, y1, x3, y3)
}

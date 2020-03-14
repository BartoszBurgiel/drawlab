package canvas

import "fmt"

// Line connects two given coordinates on the canvas
func (c *Canvas) Line(x1, y1, x2, y2 int) {
	fmt.Println("m", float64(abs(y1-y2))/float64(abs(x1-x2)), "coords:", x1, y1, x2, y2)

	f := c.lineFunction(x1, y1, x2, y2)

	// set all affected fields to c.Char
	for i := x1; i <= x2; i++ {

		c.Point(i, f(i))

	}
}

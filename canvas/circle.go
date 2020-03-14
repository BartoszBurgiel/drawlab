package canvas

import "math"

// Circle draws a circle with the center at (x|y)
// and with the radius r
func (c *Canvas) Circle(x, y, r int) {

	// factor for x since length of one height unit > length of one width unit
	factor := 2.0

	// go over all possible points of a circle
	for i := 0; i < 360; i++ {

		// x-coordinate of the point on the circle
		xCoord := int(float64(r)*math.Sin(float64(i))*factor + float64(x))

		// y-coordinate of the point on the circle
		yCoord := int(float64(r)*math.Cos(float64(i)) + float64(y))

		// add points
		c.Point(xCoord, yCoord)
	}
}

package canvas

import "math"

/*
	ALL OF THE MATHEMATICAL FUNCTIONS
*/

// lineFunction calculates the linear function
// calculating each of the coordinates of the line
// if m = +Inf no (valid) function is returned
func (c Canvas) lineFunction(x1, y1, x2, y2 int) (func(x int) int, bool) {
	m := float64(y1-y2) / float64(x1-x2)
	b := int(float64(y1) - m*float64(x1))

	return func(x int) int {
		return int(m*float64(x) + float64(b))
	}, math.IsInf(m, 0)
}

// absolute value
func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}

package canvas

/*
	ALL OF THE MATHEMATICAL FUNCTIONS
*/

// lineFunction assemble calculates the linear function
// calculating each of the coordinates of the line
func (c Canvas) lineFunction(x1, y1, x2, y2 int) func(x int) int {
	m := float64(y1-y2) / float64(x1-x2)
	b := y1

	return func(x int) int {
		return int(m*float64(x) + float64(b))
	}
}

// absolute value
func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}

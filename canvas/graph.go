package canvas

// Graph represents a graph on what data can be plotted on
type Graph struct {
	xScale int
	yScale int
	label  string
	data   DataPoints
}

// DataPoints defines what methods a data set
// has to contain in order to be plotted with drawlab
type DataPoints interface {

	// GetMax number from the data set
	GetMax() int

	// GetMin number from the data set
	GetMin() int

	// Size returns the total number of data points in the set
	Size() int

	// GetElement with the index n
	GetElement(n int) int
}

// Graph creates a graph on the canvas
//
// the width of the canvas is dependent
// of the size of the data points
func (c *Canvas) Graph(x, y, height int, data DataPoints, label string, pointString string) {

	// print the label
	c.Text(x+(data.Size()/2), y-1, label)

	// print the y-axis
	// change character to the actual line
	c.SetChar("|")
	c.Line(x, y, x, y+height)
	c.SetChar("-")

	// print the x-axis
	c.Line(x, y+height, x+data.Size(), y+height)

	c.SetChar(pointString)
	// iterate over data
	for i := 0; i < data.Size(); i++ {

		// x coordinate of the data point
		// index + offset + axis line
		xCoord := i + x + 1

		// y coordinate of the data point
		// mapped to the height
		// -> element/max + height
		yCoord := int(float64(data.GetElement(i))/float64(data.GetMax())*float64(height)) + height
		c.Point(xCoord, yCoord)
	}
	c.SetChar("-")

}

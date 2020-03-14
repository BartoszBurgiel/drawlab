package canvas

// Text prints a text onto the screen
func (c *Canvas) Text(x1, y1 int, text string) {
	currChar := c.GetChar()

	// Iterate over text and add points
	for i, t := range text {

		// Set point
		c.SetChar(string(t))
		c.Point(x1+i, y1)
	}

	c.SetChar(currChar)
}

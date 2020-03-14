package canvas

// SetChar and change the current char to the
// one provided
func (c *Canvas) SetChar(b string) {
	c.Char = b
}

// GetChar returns the current char
func (c Canvas) GetChar() string {
	return c.Char
}

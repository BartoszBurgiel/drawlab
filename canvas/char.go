package canvas

// SetChar and change the current char to the
// one provided
func (c *Canvas) SetChar(b byte) {
	c.Char = string(b)
}

// GetChar returns the current char
func (c Canvas) GetChar() byte {
	return []byte(c.Char)[0]
}

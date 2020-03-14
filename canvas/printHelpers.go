package canvas

// print vertival border
func (c Canvas) printVerticalBorder() string {
	out := ""

	// Print border
	for i := 0; i <= len(c.Fields[0])+1; i++ {
		out += "="
	}

	// line break
	out += "\n"
	return out
}

// Print the title string on the center of the
// canvas
func (c Canvas) printHeader() string {
	out := ""

	for i := 0; i < (len(c.Fields[0])/2)-len(c.title)/2; i++ {
		out += " "
	}
	out += c.title
	for i := 0; i < (len(c.Fields[0])/2)-len(c.title)/2; i++ {
		out += " "
	}
	out += "\n"
	return out
}

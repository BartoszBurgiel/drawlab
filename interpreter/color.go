package interpreter

import (
	"strings"
)

// fetch information about the color from the
// line
func (p *Program) lineToColor(line string) {

	// remove whitespace
	line = strings.ReplaceAll(line, " ", "")
	line = strings.ReplaceAll(line, "\n", "")

	// get the color
	div := strings.Split(line, "(")
	color := div[1]

	// clear color
	color = strings.ReplaceAll(color, ")", "")

	// add the function to the program
	p.Funcs = append(p.Funcs, Function{
		Name:       "color",
		Parameters: []interface{}{color}},
	)
}

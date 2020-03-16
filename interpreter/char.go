package interpreter

import (
	"fmt"
	"os"
	"strings"
)

// fetch data from the line and
// set the parameter as the current char
func (p *Program) lineToChar(line string) {
	// remove whitespace
	line = strings.ReplaceAll(line, " ", "")
	line = strings.ReplaceAll(line, "\n", "")

	// get the char
	div := strings.Split(line, "(")
	char := div[1]

	// clear char
	char = strings.ReplaceAll(char, ")", "")

	// check for length
	if len(char) != 1 || char == " " {
		fmt.Println("drawlab: char must be a single character and not be a blanc")
		os.Exit(0)
	}

	// add the function to the program
	p.Funcs = append(p.Funcs, Function{
		Name:       "char",
		Parameters: []interface{}{char}},
	)
}

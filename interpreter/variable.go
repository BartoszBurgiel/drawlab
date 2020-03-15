package interpreter

import (
	"strconv"
	"strings"
)

// Variable represents a variable
// that will be declared and called
type Variable struct {
	Name  string
	Value interface{}
}

// fetch data from a line of code and
// convert it into a variable
func (p Program) lineToVariable(line string) Variable {
	var val interface{}

	// remove the first letter of the variable
	line = line[1:]

	// divide to two parts
	div := strings.Split(line, "=")

	// get the name (and remove the whitespace)
	name := strings.ReplaceAll(div[0], " ", "")

	/*
	 determine the value
	*/

	// check if it's a string
	if strings.Contains(div[1], "\"") {
		div[1] = strings.ReplaceAll(div[1], "\n", "")
		val = strings.ReplaceAll(div[1], "\"", "")
	} else if div[1][0] == '#' || div[1][1] == '#' {

		// remove all whitespace
		div[1] = strings.ReplaceAll(div[1], " ", "")
		div[1] = strings.ReplaceAll(div[1], "\n", "")

		// if it's a variable
		val = p.getVariable(div[1][1:])

	} else {

		// replace all whitespace
		div[1] = strings.ReplaceAll(div[1], " ", "")
		div[1] = strings.ReplaceAll(div[1], "\n", "")

		n, _ := strconv.Atoi(div[1])
		val = n
	}

	return Variable{
		Name:  name,
		Value: val,
	}
}

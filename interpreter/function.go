package interpreter

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Function represents a function
// that is called
type Function struct {
	Name       string
	Parameters []interface{}
}

// pull all relevant information from a line of code
// and convert it to the function struct
func (p Program) lineToFunction(line string) Function {
	if len(line) < 3 {
		return Function{}
	}

	// divide the function and the parameter
	divided := strings.Split(line, "(")

	// name of the function at the position 0
	functionName := divided[0]

	// handle and split parameter
	params := strings.Split(divided[1], ",")

	parameter := []interface{}{}
	for _, param := range params {

		// check length of the parameter
		if len(param) < 1 {
			fmt.Println("drawlab: null parameter found at the function", functionName)
			os.Exit(0)
		}

		// remove all whitespace and closed bracket

		param = strings.ReplaceAll(param, ")", "")
		param = strings.ReplaceAll(param, "\n", "")

		// check if first byte is a space
		tempParam := param
		if param[0] == ' ' {
			tempParam = param[1:]
		}

		// check if number
		n, err := strconv.Atoi(tempParam)
		if err != nil {

			// check if it's text for the text function
			if tempParam[0] == '"' {
				// remove quotes
				parameter = append(parameter, tempParam[1:len(tempParam)-1])

				// check if it's variable
			} else if tempParam[0] == '#' {

				// remove '#'
				val := p.getVariable(tempParam[1:])
				parameter = append(parameter, val)
			}
		} else {
			parameter = append(parameter, n)
		}
	}

	return Function{
		Name:       functionName,
		Parameters: parameter,
	}
}

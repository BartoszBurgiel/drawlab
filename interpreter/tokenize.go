package interpreter

import (
	"strconv"
	"strings"
)

// Tokenize divides the code into tokens
func Tokenize(code string) []Function {
	functions := []Function{}

	// split code into lines
	splittedCode := strings.SplitAfter(code, "\n")

	for _, line := range splittedCode {
		functions = append(functions, lineToFunction(line))
	}

	return functions
}

// pull all relevant information from a line of code
// and convert it to the function struct
func lineToFunction(line string) Function {
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
				param = strings.ReplaceAll(tempParam, "\"", "")
				parameter = append(parameter, tempParam)
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

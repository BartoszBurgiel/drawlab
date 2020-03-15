package interpreter

import (
	"fmt"
	"strings"
)

// Prepare divides the code into functions
// with precalculated parameters
func Prepare(code string) []Function {

	// Define program
	prog := Program{}

	// split code into lines
	splittedCode := strings.SplitAfter(code, "\n")

	for _, line := range splittedCode {

		// only if valid length
		if len(line) > 2 {

			// examine line
			switch determineType(line) {
			case "func":
				prog.Funcs = append(prog.Funcs, prog.lineToFunction(line))
				break
			case "var":
				prog.lineToVariable(line)
			}
		}

	}
	fmt.Println(prog)
	return prog.Funcs
}

// determine which type a line belongs to
func determineType(line string) string {
	if line[0] == '#' {
		return "var"
	}
	return "func"
}

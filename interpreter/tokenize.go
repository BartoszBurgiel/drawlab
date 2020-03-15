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
	lines := strings.SplitAfter(code, "\n")

	for i := 0; i < len(lines); i++ {

		// only if valid length
		if len(lines[i]) > 2 {

			// remove whitespace before the line
			for lines[i][0] == ' ' {
				lines[i] = lines[i][1:]
			}

			// examine line
			switch determineType(lines[i]) {
			case "func":
				prog.Funcs = append(prog.Funcs, prog.lineToFunction(lines[i]))
				break
			case "var":
				prog.lineToVariable(lines[i])
				break
			case "loop":
				i += prog.lineToLoop(lines[i:])
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
	} else if line[:4] == "loop" {
		return "loop"
	}
	return "func"
}

package interpreter

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Loop represents the loop
// with its set of functions and
// loop variable
type Loop struct {

	// name of the iterable variable
	Variable string

	// Amount of increasement per iteration
	Incr int

	// End of the loop
	End int
}

// fetch the loop data from the line
// and jump over to the line where
// loop ends -> return the length of the loop
func (p *Program) lineToLoop(code []string) int {
	// cut the name from the params
	params := strings.Split(code[0], "(")

	// divide the parameter
	params = strings.Split(params[1], ",")

	// check the number of params
	if len(params) > 3 {
		fmt.Println("drawlab: too much arguments provided in the loop")
		os.Exit(0)
	} else if len(params) < 3 {
		fmt.Println("drawlab: too few arguments provided in the loop")
		os.Exit(0)
	}

	// clear the params
	for i := 0; i < len(params); i++ {
		params[i] = strings.ReplaceAll(params[i], " ", "")
	}

	// clear last -> remove {, ) and \n
	params[2] = strings.ReplaceAll(params[2], "{", "")
	params[2] = strings.ReplaceAll(params[2], ")", "")
	params[2] = strings.ReplaceAll(params[2], "\n", "")

	// convert params to numbers
	end, err := strconv.Atoi(params[1])
	if err != nil {
		fmt.Println("drawlab: second loop parameter must be a numeric value", params[1], "isnt")
		os.Exit(0)
	}

	incr, err := strconv.Atoi(params[2])
	if err != nil {
		fmt.Println("drawlab: second loop parameter must be a numeric value", params[2], "isnt")
		os.Exit(0)
	}

	// assemble the loop element
	loop := Loop{
		End:      end,
		Incr:     incr,
		Variable: params[0][1:],
	}

	// cut the loop's content from the code
	loopCode := []string{}
	loopLength := 0
	for i := 1; i < len(code); i++ {

		// if length is valid
		if len(code[i]) > 1 {
			// remove whitespace before the line
			for code[i][0] == ' ' {
				code[i] = code[i][1:]
			}

			// remove \n
			code[i] = strings.ReplaceAll(code[i], "\n", "")

			// check if end of the loop
			if code[i] == "}" {
				loopLength = i
				break
			}

			// add code[i] to loopCode
			loopCode = append(loopCode, code[i])
		}
	}

	// execute code after the loop
	// get starting value
	start := p.getVariable(loop.Variable)

	for i := start.(int); i <= loop.End+1; i += loop.Incr {

		// iterate over loopCode
		for j := 0; j < len(loopCode); j++ {

			// only if valid length
			if len(loopCode[j]) > 2 {

				// remove whitespace before the line
				for loopCode[j][0] == ' ' {
					loopCode[j] = loopCode[j][1:]
				}

				// examine line
				switch determineType(loopCode[j]) {
				case "func":
					p.Funcs = append(p.Funcs, p.lineToFunction(loopCode[j]))
					break
				case "var":
					p.lineToVariable(loopCode[j])
					break
				case "loop":
					j += p.lineToLoop(loopCode[j:])
					break
				}
			}
		}

		// increase the acutal variable
		p.setVariable(loop.Variable, i)
	}
	return loopLength
}

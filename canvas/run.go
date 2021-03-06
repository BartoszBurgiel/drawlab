package canvas

import (
	"drawlab/interpreter"
	"fmt"
	"os"
)

// RunFunctions provided by drawlab's compiler
func RunFunctions(funcs []interpreter.Function) Canvas {

	// check if init function is declared
	if funcs[0].Name != "init" {
		fmt.Println("drawlab: init function missing")
		os.Exit(0)
	}

	// first function must be the init function
	can := New(funcs[0].Parameters[0].(int), funcs[0].Parameters[1].(int), funcs[0].Parameters[2].(string))

	// iterate over functions
	// determine which function is called
	for i, fun := range funcs {
		switch fun.Name {
		case "point":

			// check for legal number of parameters
			checkArguments(fun, 2, i+1)

			can.Point(fun.Parameters[0].(int), fun.Parameters[1].(int))
			break
		case "rect":

			// check for legal number of parameters
			checkArguments(fun, 4, i+1)

			can.Rect(fun.Parameters[0].(int), fun.Parameters[1].(int), fun.Parameters[2].(int), fun.Parameters[3].(int))
			break
		case "line":

			// check for legal number of parameters
			checkArguments(fun, 4, i+1)
			can.Line(fun.Parameters[0].(int), fun.Parameters[1].(int), fun.Parameters[2].(int), fun.Parameters[3].(int))
			break
		case "circle":

			// check for legal number of parameters
			checkArguments(fun, 3, i+1)
			can.Circle(fun.Parameters[0].(int), fun.Parameters[1].(int), fun.Parameters[2].(int))
			break
		case "init":
			break
		case "text":
			// check for legal number of parameters
			checkArguments(fun, 3, i+1)
			can.Text(fun.Parameters[0].(int), fun.Parameters[1].(int), fun.Parameters[2].(string))
			break
		case "triangle":
			// check for legal number of parameters
			checkArguments(fun, 6, i+1)
			can.Triangle(fun.Parameters[0].(int), fun.Parameters[1].(int), fun.Parameters[2].(int), fun.Parameters[3].(int), fun.Parameters[4].(int), fun.Parameters[5].(int))
			break
		case "color":

			// check for legal number of parameters
			checkArguments(fun, 1, i+1)

			// determine which color to set
			switch fun.Parameters[0].(string) {
			case "red":
				can.SetColor(RED)
				break
			case "green":
				can.SetColor(GREEN)
				break
			case "blue":
				can.SetColor(BLUE)
				break
			case "purple":
				can.SetColor(PURPLE)
				break
			case "cyan":
				can.SetColor(CYAN)
				break
			case "reset":
				can.SetColor(CLEAR)
				break
			case "yellow":
				can.SetColor(YELLOW)
				break
			case "white":
				can.SetColor(WHITE)
				break
			default:
				fmt.Println("drawlab: color named", fun.Parameters[0].(string), "is not supported")
				os.Exit(0)
			}
			break
		case "char":

			// check the number of provided parameters
			checkArguments(fun, 1, i+1)

			can.SetChar([]byte(fun.Parameters[0].(string))[0])
		default:
			if fun.Name != "" {
				panic("drawlab: unsupported function: " + fun.Name)
			}
		}
	}
	return can
}

// check if the number of given parameters is legal -> equal to
// the required amount
func checkArguments(fun interpreter.Function, req, line int) {
	if len(fun.Parameters) < req {
		fmt.Println("drawlab: too few arguments provided at line: " + fmt.Sprint(line))
		os.Exit(0)
	} else if len(fun.Parameters) > req {
		fmt.Println("drawlab: too many arguments provided at line: " + fmt.Sprint(line))
		os.Exit(0)
	}
}

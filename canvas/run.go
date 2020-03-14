package canvas

import (
	"drawlab/interpreter"
	"fmt"
)

// RunFunctions provided by drawlab's compiler
func RunFunctions(funcs []interpreter.Function) Canvas {

	fmt.Println(funcs)
	// first function must be the init function
	can := New(funcs[0].Parameters[0].(int), funcs[0].Parameters[1].(int), funcs[0].Name)

	// iterate over functions
	// determine which function is called
	for _, fun := range funcs {
		switch fun.Name {
		case "point":
			can.Point(fun.Parameters[0].(int), fun.Parameters[1].(int))
			break
		case "rect":
			can.Rect(fun.Parameters[0].(int), fun.Parameters[1].(int), fun.Parameters[2].(int), fun.Parameters[3].(int))
			break
		case "line":
			can.Line(fun.Parameters[0].(int), fun.Parameters[1].(int), fun.Parameters[2].(int), fun.Parameters[3].(int))
			break
		case "circle":
			can.Circle(fun.Parameters[0].(int), fun.Parameters[1].(int), fun.Parameters[2].(int))
			break
		case "init":
			break
		case "text":
			can.Text(fun.Parameters[0].(int), fun.Parameters[1].(int), fun.Parameters[2].(string))
			break
		case "triangle":
			can.Triangle(fun.Parameters[0].(int), fun.Parameters[1].(int), fun.Parameters[2].(int), fun.Parameters[3].(int), fun.Parameters[4].(int), fun.Parameters[5].(int))
			break
		default:
			if fun.Name != "" {
				panic("drawlab: unsupported function: " + fun.Name)
			}
		}
	}
	return can
}

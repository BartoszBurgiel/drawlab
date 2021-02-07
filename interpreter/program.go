package interpreter

import (
	"fmt"
	"os"
)

// Program represents the entire program
// with all of it's functions and variables
type Program struct {
	Funcs []Function
	Vars  map[string]Variable
	Color string
	Char  string
}

// construct a new program
func newProgram() *Program {
	return &Program{
		Vars: make(map[string]Variable),
	}
}

// Find a variable with the given name and return it's
// value, if no variable with this name is found
// program exists
func (p Program) getVariable(name string) interface{} {

	if v, ok := p.Vars[name]; ok {
		return v.Value
	}

	fmt.Println("drawlab: no variable with name", name, "is declared :c")
	os.Exit(0)
	return 0
}

// set variable's value to the given one
func (p *Program) setVariable(name string, val interface{}) {

	if v, ok := p.Vars[name]; ok {
		v.Value = val
		return
	}

	fmt.Println("drawlab: no variable with name", name, "is declared :c")
	os.Exit(0)
}

// check if variable exists in the program
func (p Program) checkVariable(name string) bool {
	_, ok := p.Vars[name]
	return ok
}

// print the contents of the program
// -> all functions and all variables as well as the current color and the current char
func (p Program) printProgram() {

	// print the variables
	fmt.Println("Variables:")
	for _, v := range p.Vars {
		fmt.Println("\t", v.Name, "=", v.Value)
	}

	// print the functions
	fmt.Println("Functions:")
	for _, f := range p.Funcs {
		fmt.Println("\t", f.Name, "(", f.Parameters, ")")
	}

}

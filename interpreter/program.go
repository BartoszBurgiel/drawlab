package interpreter

import (
	"fmt"
	"os"
)

// Program represents the entire program
// with all of it's functions and variables
type Program struct {
	Funcs []Function
	Vars  []Variable
}

// Find a variable with the given name and return it's
// value, if no variable with this name is found
// program exists
func (p Program) getVariable(name string) interface{} {
	// iterate over variables
	for _, v := range p.Vars {
		if v.Name == name {
			return v.Value
		}
	}

	fmt.Println("drawlab: no variable with name", name, "is declared :c")
	os.Exit(0)
	return 0
}

// set variable's value to the given one
func (p *Program) setVariable(name string, val interface{}) {
	// iterate over variables
	for i := 0; i < len(p.Vars); i++ {
		if p.Vars[i].Name == name {
			p.Vars[i].Value = val
			return
		}

	}

	fmt.Println("drawlab: no variable with name", name, "is declared :c")
	os.Exit(0)
}

// check if variable exists in the program
func (p Program) checkVariable(name string) bool {
	for _, v := range p.Vars {
		if v.Name == name {
			return true
		}
	}
	return false
}

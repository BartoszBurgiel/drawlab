package interpreter

import (
	"fmt"
	"os"
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
// handle it -> modify the value of a variable or add new
func (p *Program) lineToVariable(line string) {

	// determine what to do
	if strings.Contains(line, "=") {
		p.declareVariable(line)
	} else {
		p.modifyVariable(line)
	}

}

// declare variable and assing a value to it
// and add it to program
func (p *Program) declareVariable(line string) {
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
	} else {

		// replace all whitespace
		div[1] = strings.ReplaceAll(div[1], " ", "")
		div[1] = strings.ReplaceAll(div[1], "\n", "")

		n, _ := strconv.Atoi(div[1])
		val = n
	}

	// check if the variable already exists
	if p.checkVariable(name) {
		p.setVariable(name, val)
		return
	}

	p.Vars = append(p.Vars, Variable{Name: name, Value: val})

}

// Change variable's value
func (p *Program) modifyVariable(line string) {

	// list of all possible operators
	operators := []byte("+-/*")

	// remove line break and variable specifier
	line = strings.ReplaceAll(line, "\n", "")
	line = strings.ReplaceAll(line, "#", "")

	// check if text
	if strings.Contains(line, "\"") {
		fmt.Println("drawlab: variable operations are allowed only on numerical variables")
		os.Exit(0)
	}

	// splitted variable
	var div []string
	var operator byte

	// search for the operator
	for _, op := range operators {
		if strings.ContainsRune(line, rune(op)) {
			operator = op
			div = strings.Split(line, string(op))
		}
	}

	// get the name of the variable
	name := div[0]

	// get the value of the variable
	val := p.getVariable(name)

	// get the value to modify
	mod := strings.ReplaceAll(div[1], " ", "")

	// convert to number
	n, err := strconv.Atoi(mod)
	if err != nil {
		fmt.Println("drawlab:", mod, "is not a numerical value")
		os.Exit(0)
	}

	// determine the operation
	switch operator {
	case '+':
		newVal := val.(int) + n
		p.setVariable(name, newVal)
		break
	case '-':
		newVal := val.(int) - n
		p.setVariable(name, newVal)
		break
	case '/':
		newVal := val.(int) / n
		p.setVariable(name, newVal)
		break
	case '*':
		newVal := val.(int) * n
		p.setVariable(name, newVal)
		break
	}

}

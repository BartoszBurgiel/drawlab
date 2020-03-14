package main

import (
	"drawlab/canvas"
	"drawlab/interpreter"
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("Welcome to drawlab")

	// read the code
	code, _ := ioutil.ReadFile("code.dl")
	f := interpreter.Tokenize(string(code))

	c := canvas.RunFunctions(f)
	c.Draw()
}

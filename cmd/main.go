package main

import (
	"drawlab/canvas"
	"drawlab/interpreter"
	"fmt"
	"io/ioutil"
	"time"
)

func main() {
	fmt.Println("Welcome to drawlab")

	// read the code
	code, _ := ioutil.ReadFile("./code.dl")
	s := time.Now()
	f := interpreter.Prepare(string(code))
	c := canvas.RunFunctions(f)
	c.Draw()

	fmt.Println("Exec time:", time.Now().Sub(s))
}

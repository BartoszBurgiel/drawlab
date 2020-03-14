package main

import (
	"drawlab/canvas"
	"fmt"
)

func main() {
	fmt.Println("Welcome to drawlab")

	can := canvas.New(70, 20, "test")

	can.Line(5, 15, 60, 5)
	can.Line(5, 10, 60, 5)
	can.Line(5, 19, 60, 5)

	can.Draw()
}

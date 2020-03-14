package main

import (
	"drawlab/canvas"
	"fmt"
)

func main() {
	fmt.Println("Welcome to drawlab")

	fmt.Println("Initializing the canvas...")
	can := canvas.New(70, 20, "test")

	fmt.Println("Drawing...")
	can.Line(5, 15, 60, 5)
	can.Line(5, 10, 60, 5)
	can.Line(5, 20, 60, 5)

	can.Draw()
}

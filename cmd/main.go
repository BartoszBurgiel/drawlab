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
	// can.Line(5, 15, 60, 5)
	// can.Line(5, 10, 60, 5)
	// can.Line(5, 20, 60, 5)

	// can.Rect(10, 10, 20, 5)
	// can.Rect(5, 5, 40, 10)
	// can.Rect(40, 2, 15, 15)

	// can.Rect(10, 10, 10, 5)
	// can.Line(10, 10, 15, 7)
	// can.Line(15, 7, 20, 10)

	can.Point(30, 10)
	can.Circle(10, 10, 5)
	can.Circle(30, 10, 10)
	can.Circle(50, 10, 5)
	can.Draw()

}

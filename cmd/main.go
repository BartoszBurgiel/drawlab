package main

import (
	"drawlab/canvas"
	"drawlab/interpreter"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to drawlab")

	fmt.Println("Initializing the canvas...")
	can := canvas.New(120, 40, "test")

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
	can.SetColor(canvas.RED)
	can.Line(20, 10, 50, 10)
	can.SetColor(canvas.GREEN)
	can.Line(20, 10, 50, 17)
	can.SetColor(canvas.CYAN)
	can.Triangle(55, 7, 60, 3, 68, 7)
	can.SetColor(canvas.PURPLE)
	can.Text(30, 7, "omg it's so cool")
	can.SetColor(canvas.CLEAR)

	can.Draw()

	file, _ := os.Create("nyy.txt")
	defer file.Close()

	file.WriteString(can.ToString())
	file.Sync()

	// read the code
	code, _ := ioutil.ReadFile("code.dl")
	f := interpreter.Tokenize(string(code))

	c := canvas.RunFunctions(f)
	c.Draw()
}

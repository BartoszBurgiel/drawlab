package canvas

import "fmt"

// Canvas represents the canvas
// the user will draw on
// it is a two dimentional slice
// of bits
type Canvas struct {
	// title of the canvas
	title string

	// Fields on which the characters will append
	Fields [][]string

	// the character that is places currently
	// as well as the color
	Char string

	// Dimentions of the canvas
	Height int
	Width  int

	// Current color the shape will be printed
	color string
}

// New returns a new canvas with
// the given size (console window)
func New(width, height int, header string) Canvas {
	c := Canvas{Char: "-", title: header, Height: height, Width: width}

	// construct the two dimensional slice
	for i := 0; i < height; i++ {

		// append new row
		c.Fields = append(c.Fields, []string{})
		for j := 0; j < width; j++ {

			// append new field
			c.Fields[i] = append(c.Fields[i], " ")
		}
	}
	return c
}

// Draw the canvas with the border
func (c Canvas) Draw() {
	// clear the screen
	fmt.Println("\033[H\033[2J")

	// set color to clear
	c.SetColor(CLEAR)

	// Print the title above the canvas
	fmt.Print(c.printHeader())

	fmt.Print(c.printVerticalBorder())

	// Iterate over the canvas' fields and print
	for _, x := range c.Fields {
		fmt.Print(c.color + "|")
		for _, y := range x {
			fmt.Print(y)
		}
		// after each row add a line break
		fmt.Println(c.color + "|")
	}

	fmt.Print(c.printVerticalBorder())
}

// ToString converts the canvas into a
// string that can be exported -> written into a file
func (c Canvas) ToString() string {
	out := ""

	out += c.printHeader()
	out += c.printVerticalBorder()

	// Iterate over the canvas' fields and add to out
	for _, x := range c.Fields {
		out += "|"
		for _, y := range x {

			// add only the valid character -> the others
			// are color specs
			out += string(y[len(y)-1])
		}
		// after each row add a line break
		out += "|\n"
	}
	out += c.printVerticalBorder()
	return out
}

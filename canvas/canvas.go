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
	field := make([][]string, height)
	for i := 0; i < height; i++ {
		field[i] = make([]string, width)
	}
	return Canvas{
		Fields: field,
		Char:   "-",
		title:  header,
		Height: height,
		Width:  width,
	}

}

// Draw the canvas with the border
func (c Canvas) Draw() {
	// clear the screen
	//fmt.Println("\033[H\033[2J")

	// set color to clear
	c.SetColor(CLEAR)

	// Print the title above the canvas
	fmt.Print(c.printHeader())

	fmt.Print(c.printVerticalBorder())

	// string that holds the entire canvas
	// so that it can be printed at once
	canvString := ""
	// Iterate over the canvas' fields and print
	for _, x := range c.Fields {
		canvString += c.color + "|"

		// temporary string to combine the
		// rows
		str := ""
		for _, y := range x {
			if y == "" {
				str += " "
			} else {

				str += y
			}
		}
		canvString += str
		// after each row add a line break
		canvString += c.color + "|\n"
	}
	fmt.Print(canvString)

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

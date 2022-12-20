package main

import "fmt"

func Area(rec Rectangle) int {
	return rec.Length * rec.Width
}
func main() {
	rect := Rectangle{
		Width:  10,
		Length: 5,
	}

	area := Area(rect)

	fmt.Println("Area:", area)
}

type Rectangle struct {
	Width  int
	Length int
}

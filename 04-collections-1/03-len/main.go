package main

import "fmt"

var colors = [5]string{"red", "green", "blue"}
var systems = []string{"linux", "macos", "windows"}
var numbers = [5]int{1, 2, 3}

func NumberOfColors() int {
	return len(colors)
}

func NumberOfSystems() int {
	return len(systems)
}
func main() {
	fmt.Println(NumberOfColors(), colors)
	fmt.Println(NumberOfSystems(), systems)
	fmt.Println(len(numbers), numbers)
}

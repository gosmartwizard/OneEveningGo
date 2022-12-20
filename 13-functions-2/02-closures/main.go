package main

import "fmt"

func WordGenerator(words []string) func() string {
	index := 0
	wl := len(words)

	return func() string {
		if index == wl {
			index = 0
		}
		pos := index
		index++
		return words[pos]
	}
}

func main() {
	continents := []string{
		"Africa",
		"Antarctica",
		"Asia",
		"Australia",
		"Europe",
		"North America",
		"South America",
	}

	generator := WordGenerator(continents)

	for i := 0; i < 10; i++ {
		fmt.Println(generator())
	}
}

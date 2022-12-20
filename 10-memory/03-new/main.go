package main

import "fmt"

var limit int = 1

func AllocateBuffer() *string {

	if limit > 3 {
		return nil
	}
	limit += 1
	buf := new(string)
	return buf
}

func main() {
	var buffers []*string

	for {
		b := AllocateBuffer()
		if b == nil {
			break
		}

		buffers = append(buffers, b)
	}

	fmt.Println("Allocated", len(buffers), "buffers")
}

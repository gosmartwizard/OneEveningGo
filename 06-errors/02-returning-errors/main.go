package main

import (
	"errors"
	"fmt"
)

func Divide(num, den float64) (float64, error) {
	if den == 0 {
		return 0, errors.New("denominator is zero")
	}

	return num / den, nil
}

func main() {
	result, err := Divide(100, 50)
	fmt.Println("Result:", result, "Error:", err)
}

package main

func Sum(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum = sum + num
	}

	return sum

}
func main() {
	_ = Sum(1, 2, 3, 4, 5)
}

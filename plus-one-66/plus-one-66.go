package main

import "fmt"

func main() {
	digis := []int{8, 9}
	fmt.Println(plusOne(digis))
}

func plusOne(digits []int) []int {

	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i] = digits[i] + 1
			return digits
		}
		digits[i] = 0

	}

	digits = append([]int{1}, digits...)

	return digits
}

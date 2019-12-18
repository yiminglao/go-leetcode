package main

import "fmt"

func main() {
	a := []int{-3, -3, -1, 2}

	fmt.Println(sortedSquares(a))
}

func sortedSquares(A []int) []int {

	index := len(A) - 1

	for index >= 1 {
		left := A[0] * A[0]
		right := A[index] * A[index]
		if left > right {
			A[0] = A[index]
			A[index] = left

		} else {
			A[index] = right
		}
		index--
	}
	A[0] = A[0] * A[0]
	return A
}

// Runtime: 524 ms, faster than 21.29% of Go online submissions for Squares of a Sorted Array.
// Memory Usage: 8.5 MB, less than 66.67% of Go online submissions for Squares of a Sorted Array.

// func sortedSquares(A []int) []int {
// 	for i, v := range A {
// 		A[i] = v * v
// 	}
// 	sort.Ints(A)
// 	return A
// }

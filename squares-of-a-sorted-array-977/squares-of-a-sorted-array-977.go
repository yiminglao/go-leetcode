package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{-3, -3, -1, 2}

	fmt.Println(sortedSquares(a))
}

// func sortedSquares(A []int) []int {

// 	for i, v := range A {
// 		A[i] = v * v
// 	}

// 	for j := 1; j < len(A); j++ {
// 		key := A[j]
// 		i := j - 1
// 		for i >= 0 && A[i] > key {
// 			A[i+1] = A[i]
// 			i--
// 		}
// 		A[i+1] = key
// 	}

// 	return A
// }

// Runtime: 524 ms, faster than 21.29% of Go online submissions for Squares of a Sorted Array.
// Memory Usage: 8.5 MB, less than 66.67% of Go online submissions for Squares of a Sorted Array.

func sortedSquares(A []int) []int {
	for i, v := range A {
		A[i] = v * v
	}
	sort.Ints(A)
	return A
}

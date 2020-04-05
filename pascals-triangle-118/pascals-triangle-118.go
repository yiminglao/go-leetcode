package main

import "fmt"

func main() {
	generate(5)
}

func generate(numRows int) [][]int {
	result := make([][]int)
	for numRows > 0 {
		fmt.Println(result)
		numRows--
	}
	return result
}

package main

import "fmt"

func main() {
	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	rotate(matrix)
	fmt.Println(matrix)
}

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < (n+1)/2; i++ {
		for j := 0; j < n/2; j++ {
			temp := matrix[n-1-j][i]
			matrix[n-1-j][i] = matrix[n-1-i][n-1-j]
			matrix[n-1-i][n-j-1] = matrix[j][n-1-i]
			matrix[j][n-1-i] = matrix[i][j]
			matrix[i][j] = temp
		}
	}

}

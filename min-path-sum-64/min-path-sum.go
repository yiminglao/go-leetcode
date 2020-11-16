package main

import "fmt"

func main() {
	grid := [][]int{{1, 3, 1}}
	fmt.Println(minPathSum(grid))
}
func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	if m == 1 {
		for i := 1; i < n; i++ {
			grid[0][i] = grid[0][i] + grid[0][i-1]
		}
	} else {
		for i := 1; i < m; i++ {
			for j := 1; j < n; j++ {
				if i == 1 {
					grid[0][j] = grid[0][j-1] + grid[0][j]
				}
				if j == 1 {
					grid[i][0] = grid[i][0] + grid[i-1][0]
				}
				grid[i][j] = min(grid[i-1][j], grid[i][j-1]) + grid[i][j]
			}
		}
	}
	return grid[m-1][n-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

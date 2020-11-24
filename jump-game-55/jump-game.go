package main

import "fmt"

func main() {
	nums := []int{3, 3, 1, 0, 4}
	fmt.Println(canJump(nums))
}

func canJump(nums []int) bool {
	maxReach := 0
	for k, v := range nums {
		if k > maxReach {
			return false
		} else if maxReach >= len(nums)-1 {
			return true
		}
		maxReach = max(maxReach, k+v)
	}
	return true
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

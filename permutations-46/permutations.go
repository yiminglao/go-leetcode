package main

import "fmt"

var result [][]int

func main() {
	nums := []int{1, 2, 3}
	permute(nums)
	fmt.Println(result)

}

func permute(nums []int) [][]int {
	result := [][]int{}
	pathNum := []int{}
	isUsed := make([]bool, len(nums))
	backtracking(nums, pathNum, isUsed)
	return result

}

func backtracking(nums, pathNum []int, isUsed []bool) {
	if len(nums) == len(pathNum) {
		temp := make([]int, len(nums))
		copy(temp, pathNum)
		result = append(result, temp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if !isUsed[i] {
			isUsed[i] = true
			pathNum = append(pathNum, nums[i])
			backtracking(nums, pathNum, isUsed)
			pathNum = pathNum[:len(pathNum)-1]
			isUsed[i] = false
		}
	}

}

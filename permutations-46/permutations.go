package main

import "fmt"

var ans [][]int

func main() {
	nums := []int{1, 2, 3}
	permute(nums)
	fmt.Println(ans)

}

func permute(nums []int) [][]int {
	ans := [][]int{}
	path := []int{}
	isUsed := make([]bool, len(nums))
	backtracking(nums, path, isUsed)
	return ans
}

func backtracking(nums, path []int, isUsed []bool) {
	if len(path) == len(nums) {
		temp := make([]int, len(nums))
		copy(temp, path)
		ans = append(ans, temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if !isUsed[i] {
			isUsed[i] = true
			path = append(path, nums[i])
			backtracking(nums, path, isUsed)
			path = path[:len(path)-1]
			isUsed[i] = false
		}
	}
}

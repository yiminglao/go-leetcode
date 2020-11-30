package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))
}

func subsets(nums []int) [][]int {
	ans := [][]int{}

	var dfs func(i int, list []int)
	dfs = func(i int, list []int) {
		if i == len(nums) {
			newlist := make([]int, len(list))
			copy(newlist, list)
			ans = append(ans, newlist)
			return
		}
		list = append(list, nums[i])
		dfs(i+1, list)
		list = list[:len(list)-1]
		dfs(i+1, list)
	}

	dfs(0, []int{})

	return ans
}

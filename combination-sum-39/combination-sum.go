package main

import (
	"fmt"
	"sort"
)

func main() {
	candidates := []int{1, 2, 3, 6, 7}
	targe := 7
	fmt.Println(combinationSum(candidates, targe))
}

func combinationSum(candidates []int, target int) [][]int {
	ans := [][]int{}
	sort.Ints(candidates)
	var dfs func([]int, int, int)
	dfs = func(comb []int, start, sum int) {
		if sum == target {
			newTemp := make([]int, len(comb))
			copy(newTemp, comb)
			ans = append(ans, newTemp)
			return
		}

		for i := start; i < len(candidates); i++ {
			if sum+candidates[i] > target {
				break
			}
			comb = append(comb, candidates[i])
			dfs(comb, i, sum+candidates[i])
			comb = comb[:len(comb)-1]
		}

	}
	dfs([]int{}, 0, 0)
	return ans
}

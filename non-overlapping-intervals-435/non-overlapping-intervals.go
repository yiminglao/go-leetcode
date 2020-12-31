package main

import (
	"fmt"
	"sort"
)

func main() {
	interval := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}
	fmt.Println(eraseOverlapIntervals(interval))
}

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) <= 1 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	ans, right := 1, intervals[0][1]
	for _, v := range intervals[1:] {
		if v[0] >= right {
			ans++
			right = v[1]
		}
	}
	return len(intervals) - ans
}

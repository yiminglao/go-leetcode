package main

import "fmt"

func main() {
	nums := []int{6, 5, 5}
	fmt.Println(majorityElement(nums))

}

// func majorityElement(nums []int) int {
// 	max := 0
// 	currNum := 0
// 	hMap := make(map[int]int)

// 	for _, v := range nums {
// 		// hMap[v] = hMap[v] + 1
// 		hMap[v]++
// 		c := hMap[v]
// 		if c > max {
// 			max = c
// 			currNum = v
// 		}
// 	}
// 	return currNum
// }

func majorityElement(nums []int) int {
	major := 0
	count := 0

	for _, v := range nums {
		if count == 0 {
			major = v
			count++
			continue
		}
		if major == v {
			count++
		} else {
			count--
		}
	}
	return major

}

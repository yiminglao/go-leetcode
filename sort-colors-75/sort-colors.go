package main

import "fmt"

func main() {
	nums := []int{2, 0, 1}
	sortColors(nums)
	fmt.Println(nums)
}

// sort Color 1
// func sortColors(nums []int) {
// 	p0, p1 := 0, 0

// 	for i, v := range nums {
// 		if v == 0 {
// 			nums[p0], nums[i] = nums[i], nums[p0]
// 			if p0 < p1 {
// 				nums[p1], nums[i] = nums[i], nums[p1]
// 			}
// 			p0++
// 			p1++
// 		}
// 		if v == 1 {
// 			nums[p1], nums[i] = nums[i], nums[p1]
// 			p1++
// 		}
// 	}
// }

func sortColors(nums []int) {
	left, right := 0, len(nums)-1

	for i := 0; i <= right; i++ {
		if nums[i] == 0 {
			nums[left], nums[i] = nums[i], nums[left]
			left++
		} else if nums[i] == 2 {
			nums[right], nums[i] = nums[i], nums[right]
			right--
			i--
		}
	}

}

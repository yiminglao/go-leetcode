package main

import "fmt"

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	fmt.Println(searchRange(nums, target))
}

func searchRange(nums []int, target int) []int {
	ans := []int{-1, -1}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			ans[0], ans[1] = mid, mid
			l, r := mid-1, mid+1
			for l >= left {
				if nums[l] == target {
					ans[0] = l
				} else {
					break
				}
				l--
			}
			for r <= right {
				if nums[r] == target {
					ans[1] = r
				} else {
					break
				}
				r++
			}
			return ans
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}

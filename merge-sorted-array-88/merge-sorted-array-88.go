package main

import "fmt"

func main() {
	// m := 3
	// n := 3
	// nums1 := []int{1, 2, 3, 0, 0, 0}
	// nums2 := []int{2, 5, 6}
	m := 0
	n := 1
	nums1 := []int{0}
	nums2 := []int{2}

	merge(nums1, m, nums2, n)
}

func merge(nums1 []int, m int, nums2 []int, n int) {

	i := m - 1
	j := n - 1
	k := m + n - 1
	for ; i >= 0 && j >= 0; k-- {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
	}

	for ; j >= 0; j-- {
		nums1[j] = nums2[j]
	}

	fmt.Println(nums1)
}

// func merge(nums1 []int, m int, nums2 []int, n int) {

// 	nums1 = append(nums1[:m], nums2[:n]...)
// 	sort.Ints(nums1)
// 	fmt.Println(nums1)
// }

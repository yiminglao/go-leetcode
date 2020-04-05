package main

func main() {
	nums := []int{4, 1, 2, 1, 2}
	singleNumber(nums)
}

func singleNumber(nums []int) int {
	t := make(map[int]int)
	result := 0
	for _, v := range nums {
		if val, ok := t[v]; ok {
			delete(t, val)
			result -= val
		} else {
			t[v] = v
			result += v
		}
	}
	return result

	// if len(nums) == 0 {
	// 	return 0
	// }
	// res := nums[0]
	// for i := 1; i < len(nums); i++ {
	// 	res = res ^ nums[i]
	// }
	// return res
}

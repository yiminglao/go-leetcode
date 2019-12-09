package main

func main() {
	nums := []int{3, 2, 2, 3, 4, 5, 3, 2, 3, 3, 3, 3}
	val := 3

	println(removeElement(nums, val))

}

func removeElement(nums []int, val int) int {
	index := 0
	arrLength := len(nums)
	for index < arrLength {
		if val == nums[index] {
			nums[index] = nums[arrLength-1]
			arrLength--
		} else {
			index++
		}
	}
	return index
}

package main

import (
	"fmt"
)

func main() {
	nums := []int{-2, 0, 3, -5, 2, -1}
	obj := Constructor(nums)
	fmt.Println(obj.SumRange(0, 2))
}

// NumArray struce
type NumArray struct {
	arraySum []int
}

//Constructor numArray
func Constructor(nums []int) NumArray {
	arr := NumArray{[]int{0}}
	for idx, v := range nums {
		arr.arraySum = append(arr.arraySum, arr.arraySum[idx]+v)
	}
	return arr
}

//SumRange get result
func (this *NumArray) SumRange(i int, j int) int {
	return this.arraySum[j+1] - this.arraySum[i]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */

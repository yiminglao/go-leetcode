package main

import (
	"fmt"
	"math"
)

func main() {
	result := int(math.Floor(math.Sqrt(9)))
	fmt.Print(result)
	fmt.Println(mySqrt(9))
}

func mySqrt(x int) int {

	left, right := 1, x
	for {
		mid := left + (right-left)/2
		if mid*mid > x {
			right = mid - 1
		} else {
			if (mid+1)*(mid+1) > x {
				return mid
			}
			left = mid + 1
		}
	}

}

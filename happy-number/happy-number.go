package main

import "fmt"

func main() {
	fmt.Println(isHappy(19))
}

func isHappy(n int) bool {
	isExist := make(map[int]bool)
	temp := n

	for temp != 1 {
		if _, ok := isExist[temp]; ok {
			return false
		}
		isExist[temp] = true

		sq := 0
		for temp > 9 {
			remain := temp % 10
			temp = (temp - remain) / 10
			sq += remain * remain
		}
		temp = temp*temp + sq
	}
	return true
}

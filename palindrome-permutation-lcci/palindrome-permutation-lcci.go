package main

import "fmt"

func main() {
	result := canPermutePalindrome("tactcoa")
	fmt.Println(result)
}

func canPermutePalindrome(s string) bool {

	var m = make(map[rune]int)
	count := 0
	for _, v := range s {
		m[v]++
		if m[v]%2 != 0 {
			count++
		} else {
			count--
		}
	}
	return count <= 1
}

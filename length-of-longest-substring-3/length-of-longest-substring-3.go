package main

import "fmt"

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	m := map[byte]int{}
	size, ans := len(s), 0
	right := 0
	for i := 0; i < size; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}
		for right < size && m[s[right]] == 0 {
			m[s[right]]++
			right++
		}
		ans = max(ans, right-i)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

package main

import "fmt"

func main() {
	s := "asdsa"
	fmt.Println(longestPalindrome(s))
}

// func lengthOfLongestSubstring(s string) int {
// 	m := map[byte]int{}
// 	size, ans := len(s), 0
// 	right := 0
// 	for i := 0; i < size; i++ {
// 		if i != 0 {
// 			delete(m, s[i-1])
// 		}
// 		for right < size && m[s[right]] == 0 {
// 			m[s[right]]++
// 			right++
// 		}
// 		ans = max(ans, right-i)
// 	}
// 	return ans
// }

// func max(x, y int) int {
// 	if x > y {
// 		return x
// 	}
// 	return y
// }

func longestPalindrome(s string) string {
	n := len(s)
	ans := ""
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for l := 0; l < n; l++ {
		for i := 0; i+l < n; i++ {
			if l == 0 {
				dp[i][i+l] = 1
			} else if l == 1 {
				if s[i] == s[i+l] {
					dp[i][i+l] = 1
				}
			} else {
				if s[i] == s[i+l] {
					dp[i][i+l] = dp[i+1][i+l-1]
				}
			}
			if dp[i][i+l] > 0 && l+1 > len(ans) {
				ans = s[i : i+l+1]
			}
		}
	}
	return ans
}

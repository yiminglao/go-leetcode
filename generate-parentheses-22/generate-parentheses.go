package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	ans := &[]string{}
	backtracking(n, n, "", ans)
	return *ans
}

func backtracking(left, right int, temp string, ans *[]string) {
	if right == 0 {
		*ans = append(*ans, temp)
		return
	}
	if left > 0 {
		backtracking(left-1, right, temp+"(", ans)
	}
	if right > left {
		backtracking(left, right-1, temp+")", ans)
	}
}

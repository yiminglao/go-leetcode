package main

import "fmt"

func main() {
	digits := "23"
	fmt.Println(letterCombinations(digits))
}

// func letterCombinations(digits string) []string {
// 	phoneMap := map[string]string{
// 		"2": "abc",
// 		"3": "def",
// 		"4": "ghi",
// 		"5": "jkl",
// 		"6": "mno",
// 		"7": "pqrs",
// 		"8": "tuv",
// 		"9": "wxyz",
// 	}
// 	var ans []string
// 	var dfs func(int, string)

// 	dfs = func(i int, path string) {
// 		if i >= len(digits) {
// 			ans = append(ans, path)
// 			return
// 		}

// 		for _, v := range phoneMap[string(digits[i])] {
// 			dfs(i+1, path+string(v))
// 		}

// 	}
// 	dfs(0, "")

// 	return ans

// }

var ans []string
var phoneMap = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	ans = []string{}
	backtracking(digits, 0, "")
	return ans

}

func backtracking(digits string, index int, combination string) {
	if index == len(digits) {
		ans = append(ans, combination)
	} else {
		letter := phoneMap[string(digits[index])]
		letterCount := len(letter)
		for i := 0; i < letterCount; i++ {
			backtracking(digits, index+1, combination+string(letter[i]))
		}
	}

}

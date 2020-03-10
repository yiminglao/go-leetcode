package main

import "strconv"

func main() {
	println(countAndSay(5))
}

func countAndSay(n int) string {
	result := "1"
	for i := 1; i < n; i++ {
		charCount := ""
		tempCount := 1
		for index, c := range result {
			currentChar := string(c)
			if index+1 >= len(result) || string(c) != string(result[index+1]) {
				charCount = charCount + strconv.Itoa(tempCount) + currentChar
				tempCount = 0
			}
			tempCount++
		}
		result = charCount
	}

	return result
}

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	str := strings.Trim("Today is a nice day", " ")

	fmt.Println(lengthOfLastWord(str))
}

func lengthOfLastWord(s string) int {
	str := strings.TrimSpace(s)
	str2 := strings.Split(str, " ")
	for i := len(str2) - 1; i >= 0; i-- {
		for _, value := range str2[i] {
			if !unicode.IsLetter(value) {
				return 0
			}
		}
		return len(str2[i])
	}
}

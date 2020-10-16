package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	fmt.Println(compressString("abbccd"))
}

func compressString(S string) string {
	if S == "" {
		return S
	}
	count := 1
	currValue := S[0]
	var result strings.Builder
	for _, v := range S[1:] {
		if byte(v) == currValue {
			count++
		} else {
			result.WriteByte(currValue)
			result.WriteString(strconv.Itoa(count))
			currValue = byte(v)
			count = 1
		}
	}
	result.WriteByte(currValue)
	result.WriteString(strconv.Itoa(count))
	if len(result.String()) >= len(S) {
		return S
	}

	return result.String()
}

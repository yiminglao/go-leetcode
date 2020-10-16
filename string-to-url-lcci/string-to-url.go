package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "Mr John Smith"
	fmt.Println(replaceSpaces(s, 13))

}

// func replaceSpaces(S string, length int) string {
// 	str := strings.Builder{}
// 	for i := 0; i < length; i++ {
// 		if S[i] == ' ' {
// 			str.WriteString("%20")
// 		} else {
// 			str.WriteString(string(S[i]))
// 		}
// 	}
// 	return str.String()
// }

func replaceSpaces(S string, length int) string {
	return strings.ReplaceAll(S[:length], " ", "%20")
}

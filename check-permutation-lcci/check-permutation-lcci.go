package main

import (
	"fmt"
)

func main() {
	s1, s2 := "abc", "bca"
	fmt.Println(CheckPermutation(s1, s2))
}

func CheckPermutation(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	if len(s1) != len(s2) {
		return false
	}

	m := map[byte]int{}
	for i := 0; i < len(s1); i++ {
		m[s1[i]]++
		m[s2[i]]--
	}

	for _, v := range m {
		if v > 0 {
			return false
		}
	}
	return true

}

//using go package
// func CheckPermutation(s1 string, s2 string) bool {
// 	ss1 := strings.Split(s1, "")
// 	ss2 := strings.Split(s2, "")
// 	sort.Strings(ss1)
// 	sort.Strings(ss2)
// 	return reflect.DeepEqual(ss1, ss2)
// }

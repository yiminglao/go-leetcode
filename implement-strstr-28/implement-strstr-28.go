package main

func main() {
	haystack := "hello"
	needle := "ll"

	println(strStr(haystack, needle))
}

func strStr(haystack string, needle string) int {
	if needle == "" || needle == haystack {
		return 0
	} else if len(needle) > len(haystack) {
		return -1
	}

	for i, v := range haystack {
		if string(needle[0]) == string(v) && len(needle) < len(haystack)-i {
			if haystack[i:len(needle)+i] == needle {
				return i
			}
		}
	}

	return -1

}

package main

import "strconv"

func main() {
	a := "11"
	b := "11111"

	println(addBinary(a, b))
}

func addBinary(a string, b string) string {

	i, _ := strconv.Atoi(a)
	j, _ := strconv.Atoi(b)
	result := strconv.Itoa(i + j)

	for i := len(result) - 1; i >= 0; i-- {
		if result[i] == '2' {
			result = "0"
		}
		break
	}
	println(result)

	return ""
}

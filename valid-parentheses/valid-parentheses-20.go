package main

func main() {
	s := "()}[]{}"
	println(isValid(s))
}

func isValid(s string) bool {
	stack := []rune{}
	runeMapper := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, v := range s {
		switch v {
		case '(', '[', '{':
			stack = append(stack, v)

		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != runeMapper[v] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

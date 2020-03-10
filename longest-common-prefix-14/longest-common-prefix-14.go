package main

func main() {
	str := []string{"flower", "flow", "flight"}


	println(longestCommonPrefix(str))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
			return ""
	} else if len(strs) == 1 {
			return strs[0]
	}


	result := strs[0]
for _, v := range strs {
	result = lcpHelper(result, v)
}
	return result
}

func lcpHelper(prefix string, str string) string {
pref := ""
for i, v := range prefix {
	if len(str) > i && string(str[i]) == string(v) {
		pref += string(v)
	}else {
		break
	}
}
return pref
}

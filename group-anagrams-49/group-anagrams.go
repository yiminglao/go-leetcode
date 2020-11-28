package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}

func groupAnagrams(strs []string) [][]string {
	ans := [][]string{}
	hmap := map[string][]string{}
	for _, v := range strs {
		c := []byte(v)
		sort.Slice(c, func(i, j int) bool {
			return c[i] < c[j]
		})
		hmap[string(c)] = append(hmap[string(c)], v)
	}
	for _, v := range hmap {
		ans = append(ans, v)
	}
	return ans
}

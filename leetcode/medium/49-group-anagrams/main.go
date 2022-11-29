package main

import (
	"fmt"
	"sort"
	"strings"
)

// https://leetcode.com/problems/group-anagrams/
func groupAnagrams(strs []string) [][]string {
	stringsMap := make(map[string][]string)
	for _, str := range strs {
		tmp := make([]string, 0)
		for _, ch := range str {
			tmp = append(tmp, string(ch))
		}
		sort.Strings(tmp)
		key := strings.Join(tmp, "#")
		stringsMap[key] = append(stringsMap[key], str)
	}

	res := make([][]string, 0)

	for _, arr := range stringsMap {
		res = append(res, arr)
	}

	return res
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	res := groupAnagrams(strs)
	fmt.Printf("%v\n", res)
}

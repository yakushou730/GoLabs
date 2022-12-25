package main

import (
	"reflect"
)

// https://leetcode.com/problems/find-all-anagrams-in-a-string/
func findAnagrams(s string, p string) []int {
	if len(p) > len(s) {
		return []int{}
	}

	dst := make(map[int]int)
	for i := 0; i < len(p); i++ {
		dst[int(p[i])] += 1
	}

	tmp := make(map[int]int)
	for i := 0; i < len(p)-1; i++ {
		tmp[int(s[i])] += 1
	}

	var ans []int
	for right := len(p) - 1; right < len(s); right++ {
		left := right - len(p) + 1
		tmp[int(s[right])] += 1

		res := reflect.DeepEqual(tmp, dst)

		if res {
			ans = append(ans, left)
		}

		if tmp[int(s[left])] > 1 {
			tmp[int(s[left])] -= 1
		} else {
			delete(tmp, int(s[left]))
		}
	}

	return ans
}

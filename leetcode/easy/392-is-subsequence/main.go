package main

// https://leetcode.com/problems/is-subsequence/
func isSubsequence(s string, t string) bool {
	i, j := 0, 0

	for j < len(t) && i < len(s) {
		if s[i] != t[j] {
			j++
		} else {
			i++
			j++
		}
	}

	return i == len(s)
}

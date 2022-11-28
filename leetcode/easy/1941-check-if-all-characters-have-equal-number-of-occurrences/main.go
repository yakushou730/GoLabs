package main

// https://leetcode.com/problems/check-if-all-characters-have-equal-number-of-occurrences/
func areOccurrencesEqual(s string) bool {
	strMap := make(map[int32]int)
	for _, c := range s {
		if _, ok := strMap[c]; ok {
			strMap[c] += 1
		} else {
			strMap[c] = 1
		}
	}

	count := strMap[int32(s[0])]
	for _, v := range strMap {
		if count != v {
			return false
		}
	}

	return true
}

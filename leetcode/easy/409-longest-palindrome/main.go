package main

// https://leetcode.com/problems/longest-palindrome/
func longestPalindrome(s string) int {
	charMap := make(map[int32]int)
	evenCount := 0

	for _, ch := range s {
		if charMap[ch] == 1 {
			delete(charMap, ch)
			evenCount++
		} else {
			charMap[ch] = 1
		}
	}

	var ans int

	if len(charMap) == 0 {
		ans = evenCount * 2
	} else {
		ans = evenCount*2 + 1
	}

	return ans
}

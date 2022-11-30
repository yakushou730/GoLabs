package main

// https://leetcode.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {
	left, ans := 0, 0
	charMap := make(map[uint8]struct{})

	for right := 0; right < len(s); right++ {
		ch := s[right]

		if _, ok := charMap[ch]; ok {
			for {
				curr := s[left]
				delete(charMap, s[left])
				left++

				if curr == ch {
					break
				}
			}
		}

		charMap[ch] = struct{}{}
		if len(charMap) > ans {
			ans = len(charMap)
		}
	}

	return ans
}

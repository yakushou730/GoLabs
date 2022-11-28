package main

// https://leetcode.com/problems/first-letter-to-appear-twice/
func repeatedCharacter(s string) byte {
	charMap := make(map[byte]int)
	for _, v := range s {
		_, ok := charMap[byte(v)]
		if !ok {
			charMap[byte(v)] = 1
			continue
		} else {
			charMap[byte(v)] += 1
		}

		if charMap[byte(v)] > 1 {
			return byte(v)
		}
	}

	return 0
}

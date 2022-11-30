package main

// https://leetcode.com/problems/check-if-the-sentence-is-pangram/
func checkIfPangram(sentence string) bool {
	charMap := make(map[byte]struct{})
	for _, c := range sentence {
		_, ok := charMap[byte(c)]
		if !ok {
			charMap[byte(c)] = struct{}{}
		}
		if len(charMap) == 26 {
			return true
		}
	}
	return false
}

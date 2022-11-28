package main

// https://leetcode.com/problems/maximum-number-of-balloons/
func maxNumberOfBalloons(text string) int {
	charMap := make(map[int32]int)
	for _, ch := range text {
		charMap[ch] += 1
	}

	ans := 0

	for charMap['b'] >= 1 && charMap['a'] >= 1 && charMap['l'] >= 2 && charMap['o'] >= 2 && charMap['n'] >= 1 {
		ans++
		charMap['b'] -= 1
		charMap['a'] -= 1
		charMap['l'] -= 2
		charMap['o'] -= 2
		charMap['n'] -= 1
		charMap['s'] -= 1
	}

	return ans
}

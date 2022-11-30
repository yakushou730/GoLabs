package main

// https://leetcode.com/problems/jewels-and-stones/
func numJewelsInStones(jewels string, stones string) int {
	charMap := make(map[int32]int)

	for _, ch := range stones {
		charMap[ch] += 1
	}

	ans := 0
	for _, ch := range jewels {
		if count, ok := charMap[ch]; ok {
			ans += count
		}
	}

	return ans
}

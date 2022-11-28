package main

// https://leetcode.com/problems/largest-unique-number/
func largestUniqueNumber(nums []int) int {
	intMap := make(map[int]int)
	for _, num := range nums {
		intMap[num] += 1
	}

	for k, v := range intMap {
		if v > 1 {
			delete(intMap, k)
		}
	}

	if len(intMap) == 0 {
		return -1
	}

	ans := 0
	for k, _ := range intMap {
		if k > ans {
			ans = k
		}
	}

	return ans
}

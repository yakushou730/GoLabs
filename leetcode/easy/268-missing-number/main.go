package main

// https://leetcode.com/problems/missing-number/
func missingNumber(nums []int) int {
	intMap := make(map[int]struct{})
	for _, num := range nums {
		intMap[num] = struct{}{}
	}

	for i := 0; i <= len(nums); i++ {
		_, ok := intMap[i]
		if !ok {
			return i
		}
	}

	return 0
}

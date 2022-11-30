package main

import "sort"

// https://leetcode.com/problems/intersection-of-multiple-arrays/
func intersection(nums [][]int) []int {
	intMap := make(map[int]int)
	for _, num := range nums {
		for _, item := range num {
			if _, ok := intMap[item]; ok {
				intMap[item] += 1
			} else {
				intMap[item] = 1
			}
		}
	}

	ans := make([]int, 0)
	for k, v := range intMap {
		if v == len(nums) {
			ans = append(ans, k)
		}
	}

	sort.Ints(ans)

	return ans
}

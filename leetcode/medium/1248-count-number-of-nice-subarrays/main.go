package main

// https://leetcode.com/problems/count-number-of-nice-subarrays/
func numberOfSubarrays(nums []int, k int) int {
	curr, ans := 0, 0
	counts := make(map[int]int)
	counts[0] = 1
	for _, num := range nums {
		// preSum
		curr += num % 2
		count, ok := counts[curr-k]
		if ok {
			ans += count
		}
		counts[curr] += 1
	}

	return ans
}

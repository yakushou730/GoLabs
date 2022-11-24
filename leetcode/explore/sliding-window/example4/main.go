package main

// https://leetcode.com/explore/interview/card/leetcodes-interview-crash-course-data-structures-and-algorithms/703/arraystrings/4502/
func findBestSubarray(nums []int, k int) int {
	curr := 0
	for i := 0; i < k; i++ {
		curr += nums[i]
	}
	ans := curr

	for i := k; i < len(nums); i++ {
		curr = curr + nums[i] - nums[i-k]
		if curr > ans {
			ans = curr
		}
	}

	return ans
}

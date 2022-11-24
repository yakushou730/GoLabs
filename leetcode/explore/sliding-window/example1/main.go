package main

// https://leetcode.com/explore/interview/card/leetcodes-interview-crash-course-data-structures-and-algorithms/703/arraystrings/4502/
func findLength(nums []int, k int) int {
	left, curr, ans := 0, 0, 0

	for right := 0; right < len(nums); right++ {
		curr += nums[right]

		for curr > k {
			curr -= nums[left]
			left++
		}

		count := right - left + 1
		if count > ans {
			ans = count
		}
	}

	return ans
}

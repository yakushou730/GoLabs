package main

// https://leetcode.com/problems/minimum-value-to-get-positive-step-by-step-sum/
func minStartValue(nums []int) int {
	min := nums[0]
	ans := min

	for i := 1; i < len(nums); i++ {
		min += nums[i]

		if min < ans {
			ans = min
		}
	}

	if ans <= 0 {
		ans = (-1 * ans) + 1
	} else {
		ans = 1
	}

	return ans
}

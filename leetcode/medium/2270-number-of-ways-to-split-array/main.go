package main

// https://leetcode.com/problems/number-of-ways-to-split-array/
func waysToSplitArray(nums []int) int {
	preSum := make([]int, len(nums))
	preSum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		preSum[i] = preSum[i-1] + nums[i]
	}

	ans := 0
	for i := 0; i < len(nums)-1; i++ {
		left := preSum[i]
		right := preSum[len(nums)-1] - preSum[i]
		if left >= right {
			ans++
		}
	}

	return ans
}

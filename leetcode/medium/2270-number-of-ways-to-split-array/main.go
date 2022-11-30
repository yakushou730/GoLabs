package main

// https://leetcode.com/problems/number-of-ways-to-split-array/
// https://leetcode.com/explore/interview/card/leetcodes-interview-crash-course-data-structures-and-algorithms/703/arraystrings/4503/
func waysToSplitArray(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	left := 0
	ans := 0
	for i := 0; i < len(nums)-1; i++ {
		left += nums[i]
		if left >= sum-left {
			ans++
		}
	}

	return ans
}

//func waysToSplitArray(nums []int) int {
//	preSum := make([]int, len(nums))
//	preSum[0] = nums[0]
//	for i := 1; i < len(nums); i++ {
//		preSum[i] = preSum[i-1] + nums[i]
//	}
//
//	ans := 0
//	for i := 0; i < len(nums)-1; i++ {
//		left := preSum[i]
//		right := preSum[len(nums)-1] - preSum[i]
//		if left >= right {
//			ans++
//		}
//	}
//
//	return ans
//}

package main

// https://leetcode.com/problems/running-sum-of-1d-array/
func runningSum(nums []int) []int {
	res := make([]int, 0, len(nums))

	for i := 0; i < len(nums); i++ {
		tmp := 0
		for j := 0; j <= i; j++ {
			tmp += nums[j]
		}
		res = append(res, tmp)
	}

	return res
}

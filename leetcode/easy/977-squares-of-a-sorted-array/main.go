package main

// https://leetcode.com/problems/squares-of-a-sorted-array/
func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	i, j := 0, len(nums)-1

	for index := j; index >= 0; index-- {
		squaredI := nums[i] * nums[i]
		squaredJ := nums[j] * nums[j]

		if squaredI > squaredJ {
			res[index] = squaredI
			i++
		} else {
			res[index] = squaredJ
			j--
		}
	}

	return res
}

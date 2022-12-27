package main

// https://leetcode.com/problems/find-pivot-index
func pivotIndex(nums []int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}

	// i = 0 的情況
	left := 0
	right := sum - nums[0]
	if left == right {
		return 0
	}

	for i := 1; i < len(nums); i++ {
		left += nums[i-1]
		right -= nums[i]
		if left == right {
			return i
		}
	}

	return -1
}

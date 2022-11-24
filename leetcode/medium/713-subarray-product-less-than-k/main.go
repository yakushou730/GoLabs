package main

// https://leetcode.com/problems/subarray-product-less-than-k/
func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}

	left, ans := 0, 0
	curr := 1

	for right := 0; right < len(nums); right++ {
		curr *= nums[right]

		for curr >= k {
			curr /= nums[left]
			left++
		}

		count := right - left + 1
		ans += count
	}

	return ans
}

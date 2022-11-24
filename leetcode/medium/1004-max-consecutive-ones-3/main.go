package main

// https://leetcode.com/problems/max-consecutive-ones-iii/
func longestOnes(nums []int, k int) int {
	left, curr, ans := 0, 0, 0

	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			curr++
		}

		for curr > k {
			if nums[left] == 0 {
				curr--
			}
			left++
		}

		count := right - left + 1
		if count > ans {
			ans = count
		}
	}

	return ans
}

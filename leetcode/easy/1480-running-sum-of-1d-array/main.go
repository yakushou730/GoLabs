package main

// https://leetcode.com/problems/running-sum-of-1d-array/
func runningSum(nums []int) []int {
	res := make([]int, len(nums))

	res[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		res[i] = res[i-1] + nums[i]
	}

	return res
}

//func runningSum(nums []int) []int {
//	res := make([]int, 0, len(nums))
//
//	for i := 0; i < len(nums); i++ {
//		tmp := 0
//		for j := 0; j <= i; j++ {
//			tmp += nums[j]
//		}
//		res = append(res, tmp)
//	}
//
//	return res
//}

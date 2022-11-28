package main

// https://leetcode.com/problems/subarray-sum-equals-k/
func subarraySum(nums []int, k int) int {
	sum, count := 0, 0
	sumMap := make(map[int]int)
	sumMap[0] = 1
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		diff := sum - k
		if currCount, ok := sumMap[diff]; ok {
			count += currCount
		}
		sumMap[sum] += 1
	}

	return count
}

//func subarraySum(nums []int, k int) int {
//	preSum := make([]int, len(nums))
//	preSum[0] = nums[0]
//	for i := 1; i < len(nums); i++ {
//		preSum[i] = preSum[i-1] + nums[i]
//	}
//
//	count := 0
//	sumMap := make(map[int]int)
//	sumMap[0] = 1
//	for i := 0; i < len(preSum); i++ {
//		diff := preSum[i] - k
//		if currCount, ok := sumMap[diff]; ok {
//			count += currCount
//		}
//		sumMap[preSum[i]] += 1
//	}
//
//	return count
//}

package main

// https://leetcode.com/problems/max-sum-of-a-pair-with-equal-sum-of-digits/
func maximumSum(nums []int) int {
	ans := -1
	sumMap := make(map[int]int)
	for _, num := range nums {
		tmp := num
		key := 0
		for tmp != 0 {
			reminder := tmp % 10
			key += reminder
			tmp = tmp / 10
		}
		if v, ok := sumMap[key]; ok {
			res := v + num
			if res > ans {
				ans = res
			}
			if num > v {
				sumMap[key] = num
			}
		} else {
			sumMap[key] = num
		}
	}

	return ans
}

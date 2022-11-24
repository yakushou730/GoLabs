package main

// https://leetcode.com/problems/two-sum/
func twoSum(nums []int, target int) []int {
	tmp := make(map[int]int)
	for i, v := range nums {
		tmp[v] = i
	}

	var first, second int

	for i, v := range nums {
		if ii, ok := tmp[target-v]; ok && i != ii {
			first = i
			second = ii
			break
		}
	}

	return []int{first, second}
}

package main

// https://leetcode.com/problems/counting-elements/
func countElements(arr []int) int {
	intMap := make(map[int]struct{})
	for _, v := range arr {
		intMap[v] = struct{}{}
	}

	var ans int
	for _, num := range arr {
		if _, ok := intMap[num+1]; ok {
			ans++
		}
	}

	return ans
}

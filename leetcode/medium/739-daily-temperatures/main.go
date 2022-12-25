package main

// https://leetcode.com/problems/daily-temperatures/
func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	var stack []int

	for i, temperature := range temperatures {
		for len(stack) != 0 && temperatures[stack[len(stack)-1]] < temperature {
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[j] = i - j
		}

		stack = append(stack, i)
	}

	return ans
}

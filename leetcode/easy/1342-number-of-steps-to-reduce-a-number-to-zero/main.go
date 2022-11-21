package main

// https://leetcode.com/problems/number-of-steps-to-reduce-a-number-to-zero/
func numberOfSteps(num int) int {
	steps := 0

	for num > 0 {
		if num%2 == 0 {
			num /= 2
		} else {
			num--
		}

		steps++
	}

	return steps
}

package main

// https://leetcode.com/problems/number-of-steps-to-reduce-a-number-to-zero/
func numberOfSteps(num int) int {
	times := 0

	for {
		if num == 0 {
			break
		}

		switch num % 2 {
		case 0:
			num /= 2
		default:
			num -= 1
		}

		times++
	}

	return times
}

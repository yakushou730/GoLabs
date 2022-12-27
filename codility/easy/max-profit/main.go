package main

import "fmt"

// https://app.codility.com/programmers/lessons/9-maximum_slice_problem/max_profit/
func Solution(A []int) int {
	if len(A) == 0 {
		return 0
	}

	min := A[0]
	maxProfit := 0

	for _, v := range A {
		if v < min {
			min = v
		} else {
			maxProfit = max(maxProfit, v-min)
		}
	}

	return maxProfit
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	A := []int{23171, 21011, 21123, 21366, 21013, 21367}
	fmt.Printf("ans = %d\n", Solution(A))
}

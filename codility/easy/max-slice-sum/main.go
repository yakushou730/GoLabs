package main

import "fmt"

// https://app.codility.com/programmers/lessons/9-maximum_slice_problem/max_slice_sum/
func Solution(A []int) int {
	if len(A) == 0 {
		return 0
	}

	globalMax := A[0]
	currMax := A[0]

	for i := 1; i < len(A); i++ {
		currMax = max(A[i], currMax+A[i])
		globalMax = max(globalMax, currMax)
	}

	return globalMax
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	A := []int{3, 2, -6, 4, 7, 8}
	fmt.Printf("ans = %d\n", Solution(A))
}

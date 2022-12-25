package main

import "sort"

// https://app.codility.com/programmers/lessons/6-sorting/max_product_of_three/
func Solution(A []int) int {
	sort.Ints(A)

	threePositiveMax := A[len(A)-1] * A[len(A)-2] * A[len(A)-3]
	twoNegativeOnePositiveMax := A[0] * A[1] * A[len(A)-1]

	if threePositiveMax > twoNegativeOnePositiveMax {
		return threePositiveMax
	} else {
		return twoNegativeOnePositiveMax
	}
}

package main

import (
	"math"
)

// https://app.codility.com/programmers/lessons/3-time_complexity/tape_equilibrium/

func Solution(A []int) int {
	sum := 0

	for _, v := range A {
		sum += v
	}

	min := math.MaxInt32

	lhs := A[0]
	rhs := sum - lhs

	for i := 1; i < len(A); i++ {
		diff := math.Abs(float64(lhs) - float64(rhs))

		if int(diff) < min {
			min = int(diff)
		}

		lhs += A[i]
		rhs -= A[i]
	}

	return int(min)
}

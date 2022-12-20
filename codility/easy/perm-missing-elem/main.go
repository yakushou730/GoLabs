package main

// https://app.codility.com/programmers/lessons/3-time_complexity/perm_missing_elem/

func Solution(A []int) int {
	length := len(A) + 1
	total := ((1 + length) * length) / 2

	for _, v := range A {
		total -= v
	}

	return total
}

package main

// https://app.codility.com/programmers/lessons/5-prefix_sums/passing_cars/
func Solution(A []int) int {
	var count int
	addBase := 0

	for _, v := range A {
		if v == 0 {
			addBase++
		} else {
			count += addBase
		}
	}

	if count > 1_000_000_000 {
		return -1
	}

	return count
}

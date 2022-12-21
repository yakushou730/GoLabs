package main

// https://app.codility.com/programmers/lessons/4-counting_elements/frog_river_one/
func Solution(X int, A []int) int {
	tmp := make(map[int]struct{})
	for i, v := range A {
		tmp[v] = struct{}{}
		if len(tmp) == X {
			return i
		}
	}

	return -1
}

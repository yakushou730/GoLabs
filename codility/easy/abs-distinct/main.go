package main

// https://app.codility.com/programmers/lessons/15-caterpillar_method/abs_distinct/
func Solution(A []int) int {
	tmp := make(map[int]struct{})
	for _, v := range A {
		tmp[v*v] = struct{}{}
	}

	return len(tmp)
}

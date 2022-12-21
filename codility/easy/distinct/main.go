package main

// https://app.codility.com/programmers/lessons/6-sorting/distinct/
func Solution(A []int) int {
	tmp := make(map[int]struct{})
	for _, v := range A {
		tmp[v] = struct{}{}
	}

	return len(tmp)
}

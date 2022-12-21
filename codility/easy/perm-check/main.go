package main

// https://app.codility.com/programmers/lessons/4-counting_elements/perm_check/
func Solution(A []int) int {
	tmp := make(map[int]struct{})
	var sum int
	for _, v := range A {
		tmp[v] = struct{}{}
		sum += v
	}

	expect := (1 + len(A)) * len(A) / 2

	if len(tmp) == len(A) && sum == expect {
		return 1
	} else {
		return 0
	}
}

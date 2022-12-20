package main

// https://app.codility.com/programmers/lessons/2-arrays/odd_occurrences_in_array/

func Solution(A []int) int {
	tmp := make(map[int]struct{})
	for _, v := range A {
		if _, ok := tmp[v]; ok {
			delete(tmp, v)
		} else {
			tmp[v] = struct{}{}
		}
	}
	var ans int
	for k, _ := range tmp {
		ans = k
	}

	return ans
}

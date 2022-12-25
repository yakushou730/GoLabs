package main

// https://app.codility.com/programmers/lessons/2-arrays/cyclic_rotation/

func Solution(A []int, K int) []int {
	if len(A) == 0 {
		return A
	}

	ans := A

	for i := 0; i < K; i++ {
		ans = append(ans[len(A)-1:], ans[:len(A)-1]...)
	}

	return ans
}

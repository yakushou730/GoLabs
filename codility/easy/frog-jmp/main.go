package main

// https://app.codility.com/programmers/lessons/3-time_complexity/frog_jmp/

func Solution(X int, Y int, D int) int {
	diff := Y - X
	step := diff / D
	mod := diff % D

	if mod != 0 {
		step++
	}

	return step
}

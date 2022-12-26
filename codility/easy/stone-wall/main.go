package main

import "fmt"

// https://app.codility.com/programmers/lessons/7-stacks_and_queues/stone_wall/
func Solution(H []int) int {
	if len(H) == 0 {
		return 0
	}

	blocks := 0
	var stack []int

	for _, v := range H {
		for len(stack) > 0 && stack[len(stack)-1] > v {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			stack = append(stack, v)
			blocks++
			continue
		}

		if stack[len(stack)-1] < v {
			stack = append(stack, v)
			blocks++
		}
	}

	return blocks
}

func main() {
	H := []int{8, 8, 5, 7, 9, 8, 7, 4, 8}
	fmt.Printf("ans = %d\n", Solution(H))
}

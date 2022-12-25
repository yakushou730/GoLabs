package main

// https://app.codility.com/programmers/lessons/7-stacks_and_queues/brackets/

func Solution(S string) int {
	var stack []rune

	for _, ch := range S {
		switch ch {
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return 0
			} else {
				stack = stack[:len(stack)-1]
			}
		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return 0
			} else {
				stack = stack[:len(stack)-1]
			}
		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return 0
			} else {
				stack = stack[:len(stack)-1]
			}
		default:
			stack = append(stack, ch)
		}
	}

	if len(stack) == 0 {
		return 1
	} else {
		return 0
	}
}

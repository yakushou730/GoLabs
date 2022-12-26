package main

import "fmt"

func Solution(A []int, B []int) int {
	if len(A) == 0 {
		return 0
	}

	survival := len(A)
	var stack []int

	for i := 0; i < len(A); i++ {
		if B[i] == 1 {
			stack = append(stack, A[i])
		} else {
			for len(stack) > 0 {
				if stack[len(stack)-1] > A[i] {
					survival--
					break
				} else if stack[len(stack)-1] < A[i] {
					survival--
					stack = stack[:len(stack)-1]
				}

			}
		}
	}

	return survival
}

func main() {
	A := []int{4, 3, 2, 1, 5}
	B := []int{0, 1, 0, 0, 0}

	fmt.Printf("ans = %d\n", Solution(A, B))
}

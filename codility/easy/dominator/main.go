package main

import "fmt"

func Solution(A []int) int {
	tmp := make(map[int][]int)
	for i, v := range A {
		tmp[v] = append(tmp[v], i)
	}

	for _, v := range tmp {
		if len(v) > len(A)/2 {
			return v[0]
		}
	}

	return -1
}

func main() {
	A := []int{3, 4, 3, 2, 3, -1, 3, 3}
	fmt.Printf("ans = %d\n", Solution(A))
}

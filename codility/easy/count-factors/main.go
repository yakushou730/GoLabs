package main

import (
	"fmt"
	"math"
)

// https://app.codility.com/programmers/lessons/10-prime_and_composite_numbers/count_factors/
func Solution(N int) int {
	if N == 0 {
		return 0
	}

	var count int
	for i := 1; i <= int(math.Sqrt(float64(N))); i++ {
		if N%i == 0 {
			count += 2
		}
	}

	if int(math.Sqrt(float64(N)))*int(math.Sqrt(float64(N))) == N {
		count--
	}

	return count
}

func main() {
	fmt.Printf("ans = %d\n", Solution(9))
}

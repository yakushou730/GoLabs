package main

import (
	"fmt"
	"math"
)

// https://app.codility.com/programmers/lessons/10-prime_and_composite_numbers/min_perimeter_rectangle/
func Solution(N int) int {
	sqrt := int(math.Floor(math.Sqrt(float64(N))))
	if sqrt*sqrt == N {
		return sqrt * 4
	}

	for i := sqrt; i > 0; i-- {
		if N%i == 0 {
			return (i + N/i) * 2
		}
	}

	return (1 + N) * 2
}

func main() {
	fmt.Printf("ans = %d\n", Solution(30))
}

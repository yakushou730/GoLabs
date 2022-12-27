package main

import "fmt"

// 最小公倍數 Least Common Multiple
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

// 最大公因數 Greatest Common Divisor
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	fmt.Printf("ans = %d\n", gcd(40, 18))
}

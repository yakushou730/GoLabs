package main

// https://app.codility.com/programmers/lessons/12-euclidean_algorithm/chocolates_by_numbers/
//func Solution(N int, M int) int {
//	tmp := make([]int, N)
//	i, count := 0, 0
//
//	for tmp[i] == 0 {
//		tmp[i] = 1
//		count++
//		i = (i + M) % N
//	}
//
//	return count
//}

func Solution(N int, M int) int {
	return N / gcd(N, M)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

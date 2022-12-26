package main

import "fmt"

// https://app.codility.com/programmers/lessons/8-leader/equi_leader/
func Solution(A []int) int {
	tmp := make(map[int]int)
	var leader, leaderCount int
	for _, v := range A {
		tmp[v] += 1

		if tmp[v] > leaderCount {
			leaderCount = tmp[v]
			leader = v
		}
	}

	leftLeaderCount, rightLeaderCount, ans := 0, leaderCount, 0

	for i, v := range A {
		if v == leader {
			leftLeaderCount++
			rightLeaderCount--
		}
		if (i+1)/2 < leftLeaderCount && (len(A)-i-1)/2 < rightLeaderCount {
			ans++
		}
	}

	return ans
}

func main() {
	A := []int{4, 3, 4, 4, 4, 2}
	fmt.Printf("ans = %d\n", Solution(A))
}

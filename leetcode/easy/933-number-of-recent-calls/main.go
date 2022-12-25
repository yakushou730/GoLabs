package main

import "fmt"

// https://leetcode.com/problems/number-of-recent-calls/
type RecentCounter struct {
	tmp []int
}

func Constructor() RecentCounter {
	return RecentCounter{tmp: make([]int, 0)}
}

func (r *RecentCounter) Ping(t int) int {
	r.tmp = append(r.tmp, t)

	for _, v := range r.tmp {
		if t-v > 3000 {
			r.tmp = r.tmp[1:]
		}
	}

	return len(r.tmp)
}

func main() {
	rc := Constructor()
	fmt.Println(rc.Ping(1))
	fmt.Println(rc.Ping(100))
	fmt.Println(rc.Ping(3001))
	fmt.Println(rc.Ping(3002))
}

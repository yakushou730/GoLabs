package main

// https://app.codility.com/programmers/lessons/7-stacks_and_queues/nesting/
func Solution(S string) int {
	tmp := make([]rune, 0)
	for _, ch := range S {
		if ch == '(' {
			tmp = append(tmp, ch)
		} else {
			if len(tmp) == 0 {
				return 0
			}
			tmp = tmp[:len(tmp)-1]
		}
	}

	if len(tmp) == 0 {
		return 1
	} else {
		return 0
	}
}

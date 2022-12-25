package main

// https://app.codility.com/programmers/lessons/1-iterations/binary_gap/
func Solution(N int) int {
	max, count := 0, 0
	once := false

	for N > 0 {
		mod := N % 2

		if mod == 1 {
			if once == false {
				once = true
			} else {
				if count > max {
					max = count
				}
			}
			count = 0
		} else {
			count++
		}

		N /= 2
	}

	return max
}

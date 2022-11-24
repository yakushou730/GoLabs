package main

// https://leetcode.com/explore/interview/card/leetcodes-interview-crash-course-data-structures-and-algorithms/703/arraystrings/4502/
func findLength(s string) int {
	left, curr, ans := 0, 0, 0

	for right := 0; right < len(s); right++ {
		if s[right] == '0' {
			curr++
		}

		for curr > 1 {
			if s[left] == '0' {
				curr--
			}
			left++
		}

		count := right - left + 1
		if count > ans {
			ans = count
		}
	}

	return ans
}

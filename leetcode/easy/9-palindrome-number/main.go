package main

// https://leetcode.com/problems/palindrome-number/
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	var tmp []int
	i, j := x, 0

	for {
		i, j = i/10, i%10

		tmp = append(tmp, j)

		if i == 0 {
			break
		}
	}

	for i, j := 0, len(tmp)-1; i < j; i++ {
		if tmp[i] != tmp[j-i] {
			return false
		}
	}

	return true
}

package main

// https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/
func removeDuplicates(s string) string {
	var stack []rune

	for _, ch := range s {
		if len(stack) == 0 {
			stack = append(stack, ch)
			continue
		}

		if ch == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, ch)
		}
	}

	return string(stack)
}

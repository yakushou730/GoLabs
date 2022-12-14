package main

// https://leetcode.com/problems/valid-parentheses/
func isValid(s string) bool {
	charMap := map[int32]int32{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := make([]int32, 0)

	for _, ch := range s {
		if len(stack) == 0 {
			stack = append(stack, ch)
			continue
		}

		last := stack[len(stack)-1]
		v, ok := charMap[ch]
		if ok {
			if last != v {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, ch)
		}
	}

	return len(stack) == 0
}

//func isValid(s string) bool {
//	stack := make([]int32, 0)
//
//	for _, ch := range s {
//		if len(stack) == 0 {
//			stack = append(stack, ch)
//			continue
//		}
//
//		switch ch {
//		case '(', '[', '{':
//			stack = append(stack, ch)
//		case ')':
//			if stack[len(stack)-1] == '(' {
//				stack = stack[:len(stack)-1]
//			} else {
//				return false
//			}
//		case ']':
//			if stack[len(stack)-1] == '[' {
//				stack = stack[:len(stack)-1]
//			} else {
//				return false
//			}
//		case '}':
//			if stack[len(stack)-1] == '{' {
//				stack = stack[:len(stack)-1]
//			} else {
//				return false
//			}
//		}
//	}
//
//	return len(stack) == 0
//}

//func isValid(s string) bool {
//	stack := make([]int32, 0)
//
//	for _, ch := range s {
//		if len(stack) == 0 {
//			stack = append(stack, ch)
//			continue
//		}
//
//		switch ch {
//		case '(', '[', '{':
//			stack = append(stack, ch)
//		case ')':
//			if stack[len(stack)-1] == '(' {
//				stack = stack[:len(stack)-1]
//			} else {
//				stack = append(stack, ch)
//			}
//		case ']':
//			if stack[len(stack)-1] == '[' {
//				stack = stack[:len(stack)-1]
//			} else {
//				stack = append(stack, ch)
//			}
//		case '}':
//			if stack[len(stack)-1] == '{' {
//				stack = stack[:len(stack)-1]
//			} else {
//				stack = append(stack, ch)
//			}
//		}
//	}
//
//	return len(stack) == 0
//}

package main

import "fmt"

// https://leetcode.com/problems/fizz-buzz/
func fizzBuzz(n int) []string {
	res := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		tmp := ""
		if i%3 == 0 {
			tmp += "Fizz"
		}
		if i%5 == 0 {
			tmp += "Buzz"
		}
		if tmp == "" {
			tmp = fmt.Sprintf("%d", i)
		}

		res = append(res, tmp)
	}

	return res
}

//func fizzBuzz(n int) []string {
//	res := make([]string, 0, n)
//	for i := 1; i <= n; i++ {
//		var tmp string
//		switch {
//		case i%15 == 0:
//			tmp = "FizzBuzz"
//		case i%3 == 0:
//			tmp = "Fizz"
//		case i%5 == 0:
//			tmp = "Buzz"
//		default:
//			tmp = fmt.Sprintf("%d", i)
//		}
//
//		res = append(res, tmp)
//	}
//
//	return res
//}

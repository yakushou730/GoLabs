package main

// https://leetcode.com/problems/richest-customer-wealth/
func maximumWealth(accounts [][]int) int {
	max := 0

	for _, account := range accounts {
		tmp := 0
		for _, bank := range account {
			tmp += bank
		}
		if tmp > max {
			max = tmp
		}
	}

	return max
}

//func maximumWealth(accounts [][]int) int {
//	m := len(accounts)
//	max := 0
//
//	for i := 0; i < m; i++ {
//		n := len(accounts[i])
//		tmp := 0
//		for j := 0; j < n; j++ {
//			//fmt.Printf("accounts[%d][%d] = %d\n", i, j, accounts[i][j])
//			tmp += accounts[i][j]
//		}
//		if tmp > max {
//			max = tmp
//		}
//	}
//
//	return max
//}

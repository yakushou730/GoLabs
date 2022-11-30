package main

import "fmt"

// https://leetcode.com/problems/equal-row-and-column-pairs/
func equalPairs(grid [][]int) int {
	rowMap := make(map[string]int)
	colMap := make(map[string]int)

	n := len(grid)

	for i := 0; i < n; i++ {
		row := grid[i]
		rowKey := ""
		colKey := ""
		for j := 0; j < n; j++ {
			rowKey += fmt.Sprintf("#%d", row[j])
			colKey += fmt.Sprintf("#%d", grid[j][i])
		}
		rowMap[rowKey] += 1
		colMap[colKey] += 1
	}

	ans := 0
	for rk, rv := range rowMap {
		cv, ok := colMap[rk]
		if !ok {
			continue
		} else {
			ans += rv * cv
		}
	}

	return ans
}

package main

// https://leetcode.com/problems/number-of-islands/
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	rows := len(grid)
	cols := len(grid[0])
	var numIslands int

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == '1' {
				numIslands++
				dfs(grid, row, col)
			}
		}
	}

	return numIslands
}

func dfs(grid [][]byte, row int, col int) {
	rows := len(grid)
	cols := len(grid[0])

	if row < 0 || col < 0 || row >= rows || col >= cols || grid[row][col] == '0' {
		return
	}

	grid[row][col] = '0'

	dfs(grid, row-1, col)
	dfs(grid, row+1, col)
	dfs(grid, row, col-1)
	dfs(grid, row, col+1)
}

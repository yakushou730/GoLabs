package main

// https://leetcode.com/problems/number-of-provinces/

func findCircleNum(isConnected [][]int) int {
	visited := make([]int, len(isConnected))
	count := 0
	for i := 0; i < len(isConnected); i++ {
		if visited[i] == 0 {
			dfs(isConnected, visited, i)
			count++
		}
	}

	return count
}

func dfs(isConnected [][]int, visited []int, i int) {
	for j := 0; j < len(isConnected); j++ {
		if isConnected[i][j] == 1 && visited[j] == 0 {
			visited[j] = 1
			dfs(isConnected, visited, j)
		}
	}
}

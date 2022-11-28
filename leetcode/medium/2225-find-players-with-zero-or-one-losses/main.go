package main

import "sort"

// https://leetcode.com/problems/find-players-with-zero-or-one-losses/
func findWinners(matches [][]int) [][]int {
	winMap := make(map[int]int)
	loseMap := make(map[int]int)

	for _, match := range matches {
		winMap[match[0]] += 1
		loseMap[match[1]] += 1
	}

	neverLose := make([]int, 0)
	oneLose := make([]int, 0)

	for w, _ := range winMap {
		if _, ok := loseMap[w]; !ok {
			neverLose = append(neverLose, w)
		}
	}

	for k, v := range loseMap {
		if v == 1 {
			oneLose = append(oneLose, k)
		}
	}

	sort.Ints(neverLose)
	sort.Ints(oneLose)

	return [][]int{neverLose, oneLose}
}

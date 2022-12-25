package main

import "math"

// https://leetcode.com/problems/count-good-nodes-in-binary-tree/

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	return dfs(root, math.MinInt)
}

func dfs(node *TreeNode, curr int) int {
	if node == nil {
		return 0
	}

	var maxSoFar int
	if curr > node.Val {
		maxSoFar = curr
	} else {
		maxSoFar = node.Val
	}

	lRes := dfs(node.Left, maxSoFar)
	rRes := dfs(node.Right, maxSoFar)

	ans := lRes + rRes
	if node.Val >= maxSoFar {
		ans++
	}

	return ans
}

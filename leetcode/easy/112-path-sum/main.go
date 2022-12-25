package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var target int = 0

func hasPathSum(root *TreeNode, targetSum int) bool {
	target = targetSum

	return dfs(root, 0)
}

func dfs(node *TreeNode, curr int) bool {
	if node == nil {
		return false
	}

	if node.Left == nil && node.Right == nil {
		return (curr + node.Val) == target
	}

	curr += node.Val

	lRes := dfs(node.Left, curr)
	rRes := dfs(node.Right, curr)

	return lRes || rRes
}

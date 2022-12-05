package main

import "fmt"

// https://leetcode.com/problems/maximum-depth-of-binary-tree/

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lMax := maxDepth(root.Left)
	rMax := maxDepth(root.Right)

	var ans int
	if lMax > rMax {
		ans = lMax
	} else {
		ans = rMax
	}

	return ans + 1
}

func main() {
	tree := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	res := maxDepth(tree)

	fmt.Println(res)
}

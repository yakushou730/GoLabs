package main

import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/validate-binary-search-tree/

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return dfs(root, math.MinInt64, math.MaxInt64)
}

func dfs(node *TreeNode, small int, large int) bool {
	if node == nil {
		return true
	}

	if node.Val <= small || large <= node.Val {
		return false
	}

	left := dfs(node.Left, small, node.Val)
	right := dfs(node.Right, node.Val, large)

	return left && right
}

func main() {
	root1 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	fmt.Printf("ans1 = %v\n", isValidBST(root1))

	root2 := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
	}

	fmt.Printf("ans2 = %v\n", isValidBST(root2))

	root3 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 2,
		},
	}

	fmt.Printf("ans3 = %v\n", isValidBST(root3))

	root4 := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
		},
		Right: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	fmt.Printf("ans4 = %v\n", isValidBST(root4))

	root5 := &TreeNode{
		Val: 2147483647,
	}

	fmt.Printf("ans5 = %v\n", isValidBST(root5))
}

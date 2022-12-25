package main

import "fmt"

// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}

	return right
}

func main() {
	node1 := &TreeNode{
		Val: 3,
	}
	node2 := &TreeNode{
		Val: 5,
	}
	node3 := &TreeNode{
		Val: 1,
	}
	node4 := &TreeNode{
		Val: 6,
	}
	node5 := &TreeNode{
		Val: 2,
	}
	node6 := &TreeNode{
		Val: 0,
	}
	node7 := &TreeNode{
		Val: 8,
	}
	node8 := &TreeNode{
		Val: 7,
	}
	node9 := &TreeNode{
		Val: 4,
	}

	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = node5
	node3.Left = node6
	node3.Right = node7
	node5.Left = node8
	node5.Right = node9

	ans := lowestCommonAncestor(node1, node2, node3)

	fmt.Println(ans.Val)
}

package main

import "fmt"

// https://leetcode.com/problems/minimum-absolute-difference-in-bst/

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	arr := dfs(root)

	if len(arr) == 1 {
		return arr[0]
	}

	min := 100_000
	for i := len(arr) - 1; i > 0; i-- {
		diff := arr[i] - arr[i-1]
		if diff < min {
			min = diff
		}
	}

	return min
}

func dfs(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}
	leftNodes := dfs(node.Left)
	rightNodes := dfs(node.Right)

	var nodes []int
	nodes = append(leftNodes, node.Val)
	nodes = append(nodes, rightNodes...)

	return nodes
}

func main() {
	root1 := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 6,
		},
	}

	fmt.Printf("ans1 = %d\n", getMinimumDifference(root1))

	root2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 0,
		},
		Right: &TreeNode{
			Val: 48,
			Left: &TreeNode{
				Val: 12,
			},
			Right: &TreeNode{
				Val: 49,
			},
		},
	}

	fmt.Printf("ans2 = %d\n", getMinimumDifference(root2))
}

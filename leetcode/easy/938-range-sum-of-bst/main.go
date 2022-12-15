package main

import (
	"fmt"
)

// https://leetcode.com/problems/range-sum-of-bst/

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// DFS
func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}

	var left, right, ans int

	if root.Val >= low && root.Val <= high {
		ans += root.Val
	}

	if root.Val < high {
		right = rangeSumBST(root.Right, low, high)
	}

	if root.Val > low {
		left = rangeSumBST(root.Left, low, high)
	}

	return ans + left + right
}

// BFS
//func rangeSumBST(root *TreeNode, low int, high int) int {
//	if root == nil {
//		return 0
//	}
//
//	queue := []*TreeNode{root}
//	var ans int
//
//	for len(queue) > 0 {
//		currLevel := queue
//
//		for _, node := range currLevel {
//			if node.Left != nil {
//				queue = append(queue, node.Left)
//			}
//			if node.Right != nil {
//				queue = append(queue, node.Right)
//			}
//
//			if node.Val >= low && node.Val <= high {
//				ans += node.Val
//			}
//
//			queue = queue[1:]
//		}
//	}
//
//	return ans
//}

func main() {
	root1 := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
		Right: &TreeNode{
			Val: 15,
			Right: &TreeNode{
				Val: 18,
			},
		},
	}

	fmt.Printf("ans1 = %d\n", rangeSumBST(root1, 7, 15))

	root2 := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 1,
				},
			},
			Right: &TreeNode{
				Val: 7,
				Left: &TreeNode{
					Val: 6,
				},
			},
		},
		Right: &TreeNode{
			Val: 15,
			Left: &TreeNode{
				Val: 13,
			},
			Right: &TreeNode{
				Val: 18,
			},
		},
	}

	fmt.Printf("ans2 = %d\n", rangeSumBST(root2, 6, 10))
}

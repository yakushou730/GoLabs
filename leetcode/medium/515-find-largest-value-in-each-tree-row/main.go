package main

import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/find-largest-value-in-each-tree-row/

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	queue := []*TreeNode{root}
	var ans []int

	for len(queue) > 0 {
		currLevel := queue
		max := math.MinInt32

		for _, node := range currLevel {
			if node.Val > max {
				max = node.Val
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			queue = queue[1:]
		}

		ans = append(ans, max)
	}

	return ans
}

func main() {
	root1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 9,
			},
		},
	}

	fmt.Printf("ans1 is %v\n", largestValues(root1))

	root2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	fmt.Printf("ans2 is %v\n", largestValues(root2))
}

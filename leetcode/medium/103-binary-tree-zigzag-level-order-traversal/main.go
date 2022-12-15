package main

import "fmt"

// https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var ans [][]int

	queue := []*TreeNode{root}
	level := 1

	for len(queue) > 0 {
		currLevel := queue

		var res []int
		if level%2 == 1 {
			for i := 0; i < len(queue); i++ {
				res = append(res, queue[i].Val)
			}
		} else {
			for i := len(queue) - 1; i >= 0; i-- {
				res = append(res, queue[i].Val)
			}
		}
		ans = append(ans, res)

		for _, node := range currLevel {
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			queue = queue[1:]
		}

		level++
	}

	return ans
}

func main() {
	root1 := &TreeNode{
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

	fmt.Printf("ans1 = %v\n", zigzagLevelOrder(root1))

	root2 := &TreeNode{
		Val: 1,
	}

	fmt.Printf("ans2 = %v\n", zigzagLevelOrder(root2))

	var root3 *TreeNode
	fmt.Printf("ans3 = %v\n", zigzagLevelOrder(root3))
}

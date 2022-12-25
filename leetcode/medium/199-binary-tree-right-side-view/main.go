package main

import "fmt"

// https://leetcode.com/problems/binary-tree-right-side-view/description/

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	var ans []int
	var queue []*TreeNode

	if root != nil {
		queue = append(queue, root)
	}

	for len(queue) > 0 {
		currLevel := queue

		//fmt.Println("level count", len(currLevel))
		//for _, node := range currLevel {
		//	fmt.Println(node.Val)
		//}

		for i, node := range currLevel {
			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			queue = queue[1:]

			if i == len(currLevel)-1 {
				ans = append(ans, node.Val)
			}
		}

	}

	return ans
}

func main() {
	fmt.Println("Example 1 ...")
	root1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 4,
			},
		},
	}

	ans1 := rightSideView(root1)
	fmt.Println("ans 1 = ", ans1)

	fmt.Println("********************")

	fmt.Println("Example 2 ...")
	root2 := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 3,
		},
	}

	ans2 := rightSideView(root2)
	fmt.Println("ans2 = ", ans2)

	fmt.Println("********************")

	fmt.Println("Example 3 ...")
	var root3 *TreeNode

	ans3 := rightSideView(root3)
	fmt.Println("ans3 = ", ans3)
}

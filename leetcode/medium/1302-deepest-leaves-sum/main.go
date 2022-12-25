package main

import "fmt"

// https://leetcode.com/problems/deepest-leaves-sum/

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//type pair struct {
//	node  *TreeNode
//	depth int
//}

// DFS
//func deepestLeavesSum(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//
//	var ans, depth, currDepth int
//
//	stack := []*pair{{node: root, depth: 0}}
//
//	for len(stack) > 0 {
//		p := stack[len(stack)-1]
//		stack = stack[:len(stack)-1]
//
//		currDepth = p.depth
//
//		if p.node.Left == nil && p.node.Right == nil {
//			if currDepth > depth {
//				ans = p.node.Val
//				depth = currDepth
//			} else if currDepth == depth {
//				ans += p.node.Val
//			}
//		} else {
//			if p.node.Right != nil {
//				stack = append(stack, &pair{
//					node:  p.node.Right,
//					depth: currDepth + 1,
//				})
//			}
//			if p.node.Left != nil {
//				stack = append(stack, &pair{
//					node:  p.node.Left,
//					depth: currDepth + 1,
//				})
//			}
//		}
//	}
//
//	return ans
//}

// BFS
func deepestLeavesSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var ans int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		currLevel := queue
		ans = 0

		for _, node := range currLevel {
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			ans += node.Val

			queue = queue[1:]
		}
	}

	return ans
}

func main() {
	root1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 7,
				},
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 6,
				Right: &TreeNode{
					Val: 8,
				},
			},
		},
	}

	fmt.Printf("ans1 = %d\n", deepestLeavesSum(root1))

	root2 := &TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 9,
				},
			},
			Right: &TreeNode{
				Val: 7,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 5,
				},
			},
		},
	}

	fmt.Printf("ans2 = %d\n", deepestLeavesSum(root2))
}

package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// https://leetcode.com/problems/populating-next-right-pointers-in-each-node-ii/

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := []*Node{root}
	level := 1

	var previousNode *Node

	for len(queue) > 0 {
		currLevelNodes := queue
		previousNode = nil

		for _, node := range currLevelNodes {
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			queue = queue[1:]

			node.Next = previousNode
			previousNode = node
		}
		level++
	}

	return root
}

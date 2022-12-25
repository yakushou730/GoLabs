package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/linked-list-cycle/submissions/
func hasCycle(head *ListNode) bool {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}

//func hasCycle(head *ListNode) bool {
//	nodeMap := make(map[*ListNode]struct{})
//	for head != nil {
//		_, ok := nodeMap[head]
//		if ok {
//			return true
//		} else {
//			nodeMap[head] = struct{}{}
//		}
//
//		head = head.Next
//
//	}
//
//	return false
//}

func main() {
	node0 := ListNode{Val: 3}
	node1 := ListNode{Val: 2}
	node2 := ListNode{Val: 0}
	node3 := ListNode{Val: -4}

	node0.Next = &node1
	node1.Next = &node2
	node2.Next = &node3
	node3.Next = &node1

	fmt.Println(hasCycle(&node0))
}

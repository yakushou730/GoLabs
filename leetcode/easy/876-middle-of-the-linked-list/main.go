package main

// https://leetcode.com/problems/middle-of-the-linked-list/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	mid := head
	end := head

	for end != nil && end.Next != nil {
		mid = mid.Next
		end = end.Next.Next
	}

	return mid
}

//func middleNode(head *ListNode) *ListNode {
//	tmp := head
//	length := 0
//	listNodes := make([]*ListNode, 0)
//
//	for tmp != nil {
//		listNodes = append(listNodes, tmp)
//		length++
//		tmp = tmp.Next
//	}
//
//	return listNodes[length/2]
//}

//func middleNode(head *ListNode) *ListNode {
//	tmp := head
//	nodeCount := 1
//
//	for tmp.Next != nil {
//		nodeCount++
//		tmp = tmp.Next
//	}
//
//	middleNode := (nodeCount / 2) + 1
//
//	tmp = head
//	nodeCount = 1
//	for tmp.Next != nil {
//		tmp = tmp.Next
//		nodeCount++
//
//		if nodeCount == middleNode {
//			break
//		}
//	}
//
//	return tmp
//}

package main

// https://leetcode.com/problems/remove-duplicates-from-sorted-list/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	curr := head

	for curr != nil && curr.Next != nil {
		if curr.Next.Val == curr.Val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}

	return head
}

//func deleteDuplicates(head *ListNode) *ListNode {
//	curr := head
//
//	for curr != nil && curr.Next != nil {
//		next := curr.Next
//
//		for next.Val == curr.Val {
//			curr.Next = next.Next
//			next = next.Next
//
//			if next == nil {
//				break
//			}
//		}
//
//		curr = curr.Next
//	}
//
//	return head
//}

func main() {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  1,
				Next: nil,
			},
		},
	}
	deleteDuplicates(list)
}

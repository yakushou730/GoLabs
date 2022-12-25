package main

import "fmt"

// https://leetcode.com/problems/reverse-linked-list/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		nextTemp := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextTemp
	}
	return prev
}

//func reverseList(head *ListNode) *ListNode {
//	if head == nil {
//		return nil
//	}
//
//	var ans, prev *ListNode
//	for head != nil {
//		ans = &ListNode{
//			Val:  head.Val,
//			Next: prev,
//		}
//
//		prev = ans
//		head = head.Next
//	}
//
//	return ans
//}

func main() {
	head1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	ans1 := reverseList(head1)

	for ans1 != nil {
		fmt.Println(ans1.Val)
		ans1 = ans1.Next
	}

}

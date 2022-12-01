package main

import "fmt"

// https://leetcode.com/problems/swap-nodes-in-pairs/
// todo: needs to review the solution
// https://leetcode.com/explore/interview/card/leetcodes-interview-crash-course-data-structures-and-algorithms/704/linked-lists/4600/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := head.Next
	var prev *ListNode = nil

	for head != nil && head.Next != nil {
		if prev != nil {
			prev.Next = head.Next
		}
		prev = head

		nextNode := head.Next.Next
		head.Next.Next = head

		head.Next = nextNode
		head = nextNode
	}

	return dummy
}

func main() {
	q := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}

	ans := swapPairs(q)

	fmt.Println(ans)
}

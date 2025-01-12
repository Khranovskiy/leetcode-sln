package main

import "fmt"

func reverseList(head *ListNode) *ListNode {
	var (
		curr           = head
		prev *ListNode = nil
	)
	for curr != nil {
		next := curr.Next
		curr.Next = prev

		prev = curr
		curr = next
	}
	return prev
}

func main() {
	head := &ListNode{
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
	rev := reverseList(head)
	fmt.Println(rev.Val)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

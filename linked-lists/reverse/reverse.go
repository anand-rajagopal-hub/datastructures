package reverse

import (
	_ "fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	var current *ListNode

	for ; head != nil; head = head.Next {
		if current == nil {
			current = &ListNode{
				Val:  head.Val,
				Next: nil,
			}
		} else {
			n := &ListNode{
				Val:  head.Val,
				Next: current,
			}
			current = n
		}
	}
	return current
	//return reverse(head, nil)
	// return head
}

func reverse(head, previous *ListNode) *ListNode {

	if head == nil {
		return nil
	}
	if head.Next == nil {
		head.Next = previous
		return head
	}
	t := reverse(head.Next, head)
	head.Next = previous
	return t
}

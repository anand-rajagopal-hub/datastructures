package main

import (
	"fmt"

	"github.com/anand-rajagopal-hub/datastructures/linked-lists/reverse"
)

func add(numbers ...int) *reverse.ListNode {
	var head, current *reverse.ListNode
	for i, v := range numbers {
		if i == 0 {
			head = &reverse.ListNode{
				Val: v,
			}
			current = head
		} else {
			current.Next = &reverse.ListNode{
				Val: v,
			}
			current = current.Next
		}
	}
	return head
}

func printList(node *reverse.ListNode) {
	for ; node != nil; node = node.Next {
		fmt.Printf("%d ", node.Val)
	}
	fmt.Println()
}

func main() {
	printList(reverse.ReverseList(add(5, 2, 1, 7, 9)))
}

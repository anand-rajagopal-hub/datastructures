package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result *ListNode
	var current *ListNode
	carryOver := 0
	for l1 != nil && l2 != nil {
		t := l1.Val + l2.Val + carryOver

		if result == nil {
			result = &ListNode{}
			current = result
		} else {
			current.Next = &ListNode{}
			current = current.Next
		}

		if t > 9 {
			carryOver = t / 10
			current.Val = t % 10
		} else {
			current.Val = t
			carryOver = 0
		}
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		t := l1.Val + carryOver

		if t > 9 {
			carryOver = t / 10
			t = t % 10
		} else {
			carryOver = 0
		}
		current.Next = &ListNode{
			Val:  t,
			Next: nil,
		}

		current = current.Next

		l1 = l1.Next
	}

	for l2 != nil {
		t := l2.Val + carryOver
		if t > 9 {
			carryOver = t / 10
			t = t % 10
		} else {
			carryOver = 0
		}
		current.Next = &ListNode{
			Val:  t,
			Next: nil,
		}
		current = current.Next
		l2 = l2.Next
	}

	if carryOver > 0 {
		current.Next = &ListNode{
			Val: carryOver,
		}
	}

	return result
}

func makeList(n []int) *ListNode {
	var result *ListNode
	var current *ListNode
	for _, v := range n {
		if result == nil {
			result = &ListNode{
				Val:  v,
				Next: nil,
			}
			current = result
		} else {
			current.Next = &ListNode{
				Val:  v,
				Next: nil,
			}
			current = current.Next
		}
	}
	return result
}

func printList(l *ListNode) {
	for l != nil {
		fmt.Printf("%d ", l.Val)
		l = l.Next
	}
	fmt.Println()
}

func main() {
	l1 := makeList([]int{2, 4, 3})
	l2 := makeList([]int{5, 6, 4})
	printList(l1)
	printList(l2)
	fmt.Println("-----------------")
	printList(addTwoNumbers(l1, l2))
	fmt.Println("-----------------")

	l1 = makeList([]int{0})
	l2 = makeList([]int{0})
	printList(l1)
	printList(l2)
	fmt.Println("-----------------")
	printList(addTwoNumbers(l1, l2))
	fmt.Println("-----------------")

	l1 = makeList([]int{9, 9, 9, 9, 9, 9, 9})
	l2 = makeList([]int{9, 9, 9, 9})
	printList(l1)
	printList(l2)
	fmt.Println("-----------------")
	printList(addTwoNumbers(l1, l2))
	fmt.Println("-----------------")

	l1 = makeList([]int{2, 4, 9})
	l2 = makeList([]int{5, 6, 4, 9})
	printList(l1)
	printList(l2)
	fmt.Println("-----------------")
	printList(addTwoNumbers(l1, l2))
	fmt.Println("-----------------")
}

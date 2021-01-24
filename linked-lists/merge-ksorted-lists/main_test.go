package main

import (
	"fmt"
	"testing"
)

func makeList(nums ...int) *ListNode {
	var result, current *ListNode
	for _, item := range nums {
		if result == nil {
			result = &ListNode{
				Val:  item,
				Next: nil,
			}
			current = result
		} else {
			nn := &ListNode{
				Val:  item,
				Next: nil,
			}
			current.Next = nn
			current = current.Next
		}
	}
	printList(result)

	return result
}

func printList(l *ListNode) {
	for temp := l; temp != nil; temp = temp.Next {
		fmt.Printf("%d-->", temp.Val)
	}
	fmt.Println()
}

func equal(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	if l1 == nil && l2 == nil {
		return true
	}
	return false
}
func Test(t *testing.T) {
	testCases := []struct {
		lists    []*ListNode
		expected *ListNode
	}{
		{
			lists:    []*ListNode{makeList(1, 4, 5), makeList(1, 3, 4), makeList(2, 6)},
			expected: makeList(1, 1, 2, 3, 4, 4, 5, 6),
		},
	}
	for i, tC := range testCases {
		desc := fmt.Sprintf("Test-%d", i)
		t.Run(desc, func(t *testing.T) {
			actual := mergeKLists(tC.lists)
			printList(actual)
			if !equal(tC.expected, actual) {
				t.Errorf("Expected = %v, Actual = %v", tC.expected, actual)
			}
		})
	}
}

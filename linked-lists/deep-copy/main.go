package main

import (
	"fmt"
)

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	var result *Node
	var current *Node

	index := make(map[*Node]int)
	for i, temp := 0, head; temp != nil; temp = temp.Next {
		index[temp] = i
		i++
	}

	resultIndex := make([]*Node, len(index))

	for j := 0; head != nil; head = head.Next {
		if result == nil {
			result = &Node{
				Val: head.Val,
			}

			resultIndex[j] = result
			current = result
		} else {
			// A node might have already been made
			n, ok := index[head]
			if ok {
				fn := resultIndex[n]
				if fn != nil {
					current.Next = fn
				} else {
					current.Next = &Node{
						Val: head.Val,
					}
					resultIndex[j] = current.Next
				}
				current = current.Next
			}
		}
		if head.Random != nil {
			n, ok := index[head.Random]

			if ok {
				if resultIndex[n] != nil { // previously seen
					current.Random = resultIndex[n]
				} else {
					// pointing to a node we haven't gotten to yet
					// just create a node
					futureNode := &Node{
						Val:    head.Random.Val,
						Next:   nil,
						Random: nil,
					}
					current.Random = futureNode
					resultIndex[n] = futureNode
				}
			}
		}

		j++
	}

	return result
}

func printList(node *Node) {
	index := make(map[*Node]int)
	for i, temp := 0, node; temp != nil; temp = temp.Next {
		index[temp] = i
		i++
	}

	fmt.Println()

	for node != nil {
		if node.Random != nil {
			_, ok := index[node.Random]
			if ok {
				fmt.Printf("[%d %d %d]  ", node.Val, index[node.Random], node.Random.Val)
			} else {
				fmt.Printf("[%d %s %d]  ", node.Val, "not found", node.Random.Val)
			}
		} else {
			fmt.Printf("[%d nil] ", node.Val)
		}
		node = node.Next
	}
	fmt.Println()
}

func main() {
	var input *Node

	zero := &Node{
		Val: 7,
	}

	one := &Node{
		Val: 13,
	}

	two := &Node{
		Val: 11,
	}

	three := &Node{
		Val: 10,
	}

	four := &Node{
		Val: 1,
	}

	zero.Next = one
	one.Next = two
	two.Next = three
	three.Next = four

	zero.Random = nil
	one.Random = zero
	two.Random = four
	three.Random = two
	four.Random = zero

	input = zero

	printList(input)
	printList(copyRandomList(input))

	zero = &Node{
		Val: 1,
	}

	one = &Node{
		Val: 2,
	}

	zero.Next = one
	zero.Random = one
	one.Random = one

	input = zero

	printList(input)
	printList(copyRandomList(input))

	// var zero, one, two, three *Node

	zero = &Node{
		Val: 1,
	}

	one = &Node{
		Val: 2,
	}

	two = &Node{
		Val: 3,
	}

	three = &Node{
		Val: 4,
	}

	zero.Next = one
	one.Next = two
	two.Next = three

	zero.Random = two
	one.Random = two
	two.Random = nil
	three.Random = nil

	input = zero
	printList(input)
	printList(copyRandomList(input))

}

package main

import (
	"fmt"
)

// Node is a fundamental datastructure for a tree
type Node struct {
	value int
	left  *Node
	right *Node
}

func (n *Node) isLeaf() bool {
	return (n.left == nil && n.right == nil)
}

func insert(node *Node, value int) *Node {
	if node == nil {
		node = &Node{
			value: value,
		}
		return node
	}
	if value <= node.value {
		node.left = insert(node.left, value)
		return node
	}
	node.right = insert(node.right, value)
	return node
}

func printAscendingOrder(node *Node) {
	if node.left != nil {
		printAscendingOrder(node.left)
	}
	if node != nil {
		fmt.Printf("%d ", node.value)
	}
	if node.right != nil {
		printAscendingOrder(node.right)
	}

}

func succ(node *Node, value int) *Node {
	return successor(node, value, node)
}

func successor(node *Node, value int, potentialSuccessor *Node) *Node {
	if value == node.value && node.right == nil {
		if potentialSuccessor.value <= node.value {
			return nil
		}
		return potentialSuccessor
	}
	if value < node.value && node.left != nil {
		potentialSuccessor = node
		return successor(node.left, value, potentialSuccessor)
	}
	if value >= node.value && node.right != nil {
		return successor(node.right, value, potentialSuccessor)
	}
	return node
}

func main() {
	root := insert(nil, 6)
	insert(root, 3)
	insert(root, 4)
	insert(root, 9)
	insert(root, 7)
	insert(root, 5)
	insert(root, 2)
	insert(root, 12)
	insert(root, 0)
	insert(root, 8)
	insert(root, 14)
	insert(root, 13)
	insert(root, 1)
	printAscendingOrder(root)

	fmt.Println("finding successor - new way")
	fmt.Println("0 - ", succ(root, 0))
	fmt.Println("1 - ", succ(root, 1))
	fmt.Println("2 - ", succ(root, 2))
	fmt.Println("3 - ", succ(root, 3))
	fmt.Println("4 - ", succ(root, 4))
	fmt.Println("5 - ", succ(root, 5))
	fmt.Println("6 - ", succ(root, 6))
	fmt.Println("7 - ", succ(root, 7))
	fmt.Println("8 - ", succ(root, 8))
	fmt.Println("9 - ", succ(root, 9))
	fmt.Println("12 - ", succ(root, 12))
	fmt.Println("13 - ", succ(root, 13))
	fmt.Println("14 - ", succ(root, 14))
}

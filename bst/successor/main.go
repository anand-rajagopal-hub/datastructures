package main

import "fmt"

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

func findSuccessor(node *Node, value int) int {
	if node == nil {
		return -1
	}
	var parent *Node
	if value < node.value {
		parent = locateLeft(node, value)
	} else {
		parent = locate(node, value)
	}
	successor := successor(parent, value)
	if value == successor {
		return parent.value
	}
	return successor
}

func locateLeft(node *Node, value int) *Node {
	if node.left == nil || value >= node.left.value {
		return node
	}
	return locateLeft(node.left, value)
}

func locate(node *Node, value int) *Node {
	if node.right == nil || value < node.value {
		return node
	}
	// if node
	return locate(node.right, value)
}

func successor(node *Node, value int) int {
	fmt.Println("node ", node)
	if value < node.value && node.left != nil {
		return successor(node.left, value)
	} else if value >= node.value && node.right != nil {
		return successor(node.right, value)
	}
	return node.value
}

// 	    6
//    /   \
//   4     9
//  / \   / \
// 2   5 7  12
///       \   \
//0        8   14
// \           /
//  1         13
func main() {
	root := insert(nil, 6)
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

	fmt.Println("finding successor")
	fmt.Println("0 - ", findSuccessor(root, 0))
	fmt.Println("2 - ", findSuccessor(root, 2))
	fmt.Println("4 - ", findSuccessor(root, 4))
	fmt.Println("5 - ", findSuccessor(root, 5))
	fmt.Println("6 - ", findSuccessor(root, 6))
	fmt.Println("7 - ", findSuccessor(root, 7))
	fmt.Println("8 - ", findSuccessor(root, 8))
	fmt.Println("9 - ", findSuccessor(root, 9))
	fmt.Println("12 - ", findSuccessor(root, 12))
	fmt.Println("13 - ", findSuccessor(root, 13))
	fmt.Println("14 - ", findSuccessor(root, 14))

}

package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func symmetric(left, right *TreeNode) bool {
	//if both the nodes are nil, that's OK
	if left == nil && right == nil {
		return true
	}
	// if one of them is nil, then it's not symmetric
	if left == nil || right == nil {
		return false
	}

	// compare the values of the node and recurse
	return left.Val == right.Val && symmetric(left.Left, right.Right) && symmetric(left.Right, right.Left)
}

func isSymmetric(root *TreeNode) bool {
	return symmetric(root.Left, root.Right)
}

func main() {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
	}
	fmt.Println(isSymmetric(tree))

	tree = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
			},
		},
	}
	fmt.Println(isSymmetric(tree))
}

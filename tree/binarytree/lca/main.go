package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	var lc, rc *TreeNode
	if root.Left != nil {
		lc = lowestCommonAncestor(root.Left, p, q)
	}
	if root.Right != nil {
		rc = lowestCommonAncestor(root.Right, p, q)
	}
	if lc != nil && rc != nil {
		return root
	}
	if lc != nil {
		return lc
	}
	if rc != nil {
		return rc
	}
	return nil
}

func makeNode(val int) *TreeNode {
	return &TreeNode{
		Val: val,
	}
}

func main() {
	three := makeNode(3)
	five := makeNode(5)
	one := makeNode(1)
	six := makeNode(6)
	two := makeNode(2)
	zero := makeNode(0)
	eight := makeNode(8)
	seven := makeNode(7)
	four := makeNode(4)

	three.Left = five
	three.Right = one
	five.Left = six
	five.Right = two
	two.Left = seven
	two.Right = four
	one.Left = zero
	one.Right = eight

	root := three

	fmt.Println(lowestCommonAncestor(root, five, four))
	fmt.Println(lowestCommonAncestor(root, five, one))
	fmt.Println(lowestCommonAncestor(root, seven, four))
}

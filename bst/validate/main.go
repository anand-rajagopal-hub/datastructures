package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

var previous *TreeNode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTree(value int) *TreeNode {
	return &TreeNode{
		Val: value,
	}
}

func (t *TreeNode) insert(value int) *TreeNode {
	if value < t.Val {
		if t.Left != nil {
			t.Left.insert(value)
		} else {
			t.Left = &TreeNode{
				Val: value,
			}
		}
	}
	if value >= t.Val {
		if t.Right != nil {
			t.Right.insert(value)
		} else {
			t.Right = &TreeNode{
				Val: value,
			}
		}
	}
	return t
}

func insertSimple(node *TreeNode, value int) *TreeNode {
	if node == nil {
		return &TreeNode{
			Val: value,
		}
	}
	if value < node.Val {
		node.Left = insertSimple(node.Left, value)
	} else if value > node.Val {
		node.Right = insertSimple(node.Right, value)
	}

	return node
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return false
	}
	previous = nil
	return valid(root)
}

func valid(node *TreeNode) bool {

	if node.Left != nil {
		if !valid(node.Left) {
			return false
		}
	}

	if previous != nil && node.Val <= previous.Val {
		return false
	}

	previous = node
	if node.Right != nil {
		return valid(node.Right)
	}
	return true
}

func PrintTree(root *TreeNode) {
	if root.Left != nil {
		PrintTree(root.Left)
	}

	fmt.Printf("%d ", root.Val)
	if root.Right != nil {
		PrintTree(root.Right)
	}
}

func visualize(root *TreeNode, fn string) {
	b := &bytes.Buffer{}
	b.Write([]byte("digraph tree{\n"))
	b.Write([]byte("node [shape=Mrecord];\n"))
	draw(root, b)
	b.Write([]byte("}"))
	err := ioutil.WriteFile(fmt.Sprintf("%s.dot", fn), b.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
}

func draw(root *TreeNode, buf *bytes.Buffer) {
	buf.Write([]byte(fmt.Sprintf("%d [label=\" {<f0> %d}  \"];\n", root.Val, root.Val)))
	if root.Left != nil {
		buf.Write([]byte(fmt.Sprintf("%d [label=\" {<f0> %d} \"];\n", root.Val, root.Val)))
		buf.Write([]byte(fmt.Sprintf("%d:f0 -> %d:f0\n", root.Val, root.Left.Val)))
		draw(root.Left, buf)
	} else {
		buf.Write([]byte(fmt.Sprintf("%d [label=\"{<f0> %d} \"];\n", root.Val, root.Val)))

		buf.Write([]byte(fmt.Sprintf("lnil%d [label=\"{<f0> nil}\"];\n", root.Val)))
		buf.Write([]byte(fmt.Sprintf("%d:f0 -> lnil%d:f0\n", root.Val, root.Val)))
	}
	if root.Right != nil {
		buf.Write([]byte(fmt.Sprintf("%d [label=\"{<f0> %d} \"];\n", root.Val, root.Val)))
		buf.Write([]byte(fmt.Sprintf("%d:f0 -> %d:f0\n", root.Val, root.Right.Val)))
		draw(root.Right, buf)
	} else {
		buf.Write([]byte(fmt.Sprintf("%d [label=\"{<f0> %d} \"];\n", root.Val, root.Val)))

		buf.Write([]byte(fmt.Sprintf("rnil%d [label=\"{<f0> nil}\"];\n", root.Val)))
		buf.Write([]byte(fmt.Sprintf("%d:f0 -> rnil%d:f0\n", root.Val, root.Val)))
	}
}

func main() {
	root := NewTree(5).insert(1).insert(4).insert(3).insert(6)

	visualize(root, "tree-data")

	root = insertSimple(nil, 5)
	root = insertSimple(root, 1)
	root = insertSimple(root, 4)
	root = insertSimple(root, 3)
	root = insertSimple(root, 6)

	// PrintTree(root)
	visualize(root, "simplified-insert")

	badTree := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
		},
		Right: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	visualize(badTree, "bad-tree-data")

	fmt.Println(isValidBST(badTree))
}

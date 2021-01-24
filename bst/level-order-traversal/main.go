package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Node  *TreeNode
	Next  *ListNode
	Level int
}

type Queue struct {
	head  *ListNode
	tail  *ListNode
	count int
}

func (q *Queue) empty() bool {
	return q.count <= 0
}

func (q *Queue) add(t *TreeNode, level int) {
	if q.head == nil {
		q.head = &ListNode{
			Node:  t,
			Next:  nil,
			Level: level,
		}
		q.tail = q.head
	} else {
		q.tail.Next = &ListNode{
			Node:  t,
			Next:  nil,
			Level: level,
		}
		q.tail = q.tail.Next
	}
	q.count++
}

func (q *Queue) remove() (*TreeNode, int) {
	node := q.head.Node
	level := q.head.Level
	q.head = q.head.Next
	q.count--
	return node, level
}

func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	q := &Queue{}
	level := 0
	q.add(root, level)

	for !q.empty() {
		node, level := q.remove()
		if len(result) <= level {
			result = append(result, make([]int, 0))
		}
		result[level] = append(result[level], node.Val)
		level++
		if node.Left != nil {
			q.add(node.Left, level)
		}
		if node.Right != nil {
			q.add(node.Right, level)
		}

	}

	return result
}

func main() {

}

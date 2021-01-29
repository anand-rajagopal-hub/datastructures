package main

import (
	"fmt"
	"testing"
)

func makeTree(nums ...int) *TreeNode {
	root := &TreeNode{
		Val: nums[0],
	}
	for i := 1; i < len(nums); i++ {
		insert(nums[i], root)
	}
	return root
}

func insert(num int, root *TreeNode) {
	if num <= root.Val {
		if root.Left == nil {
			root.Left = &TreeNode{
				Val: num,
			}
			return
		}
		insert(num, root.Left)
	} else {
		if root.Right == nil {
			root.Right = &TreeNode{
				Val: num,
			}
			return
		}
		insert(num, root.Right)
	}
}

func Test(t *testing.T) {
	testCases := []struct {
		tree *TreeNode
	}{
		{
			tree: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   9,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
		},
	}
	for i, tC := range testCases {
		desc := fmt.Sprintf("Test-%d", i)
		t.Run(desc, func(t *testing.T) {
			fmt.Println(levelOrder(tC.tree))
		})
	}
}

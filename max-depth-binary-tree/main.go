package main

import (
	"fmt"
)

// TreeNode is binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	c := &TreeNode{15, nil, nil}
	b := &TreeNode{20, c, nil}
	f := &TreeNode{9, nil, nil}
	a := &TreeNode{3, b, f}
	fmt.Println(maxDepth(a))
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	depth := 1

	depth = max(depth, maxDepth(root.Left)+1)
	depth = max(depth, maxDepth(root.Right)+1)

	return depth
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

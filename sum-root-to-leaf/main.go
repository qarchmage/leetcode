package main

import (
	"fmt"
)

// TreeNode is an implementation of a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	h := &TreeNode{1, nil, nil}
	g := &TreeNode{0, nil, nil}
	f := &TreeNode{1, nil, nil}
	d := &TreeNode{0, nil, nil}
	c := &TreeNode{1, g, h}
	b := &TreeNode{0, d, f}
	a := &TreeNode{1, b, c}
	fmt.Println(sumRootToLeaf(a))
}

func sumRootToLeaf(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(root *TreeNode, cur int) int {
	if root == nil {
		return 0
	}
	cur = cur<<1 | root.Val

	fmt.Println(cur)
	if root.Left == nil && root.Right == nil {
		return cur
	}
	return dfs(root.Left, cur) + dfs(root.Right, cur)
}

func printBinaryTree(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Println(node.Val)
	printBinaryTree(node.Left)
	printBinaryTree(node.Right)

}

package main

import "fmt"

// TreeNode is a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	g := &TreeNode{1, nil, nil}
	f := &TreeNode{3, nil, nil}
	e := &TreeNode{2, f, g}
	d := &TreeNode{6, nil, nil}
	c := &TreeNode{9, nil, nil}
	b := &TreeNode{7, c, d}
	a := &TreeNode{4, b, e}

	fmt.Println(invertTree(a))

	printPreOrder(a)
}

// swaps the treenode than calls it's children
func invertTree(root *TreeNode) *TreeNode {
	if root != nil {
		root.Left, root.Right = root.Right, root.Left
		invertTree(root.Left)
		invertTree(root.Right)
	}
	return root
}

// print the tree dfs
func printPreOrder(n *TreeNode) {
	if n == nil {
		return
	}
	fmt.Printf("%d ", n.Val)
	printPreOrder(n.Left)
	printPreOrder(n.Right)
}

package main

import "fmt"

func main() {
	a := &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}
	fmt.Println(searchBST(a, 1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	a := make([]*TreeNode, 0)

	a = walk(root, a)

	for i := range a {
		if val == a[i].Val {
			return a[i]
		}
	}

	return nil
}

func walk(root *TreeNode, dings []*TreeNode) []*TreeNode {
	if root == nil {
		return dings
	}
	dings = walk(root.Left, dings)
	dings = append(dings, root)
	dings = walk(root.Right, dings)

	return dings

}

func searchBSTFaster(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}

	// check left side
	if ret := searchBST(root.Left, val); ret != nil {
		return ret
	}

	// check right side
	if ret := searchBST(root.Right, val); ret != nil {
		return ret
	}

	return nil
}

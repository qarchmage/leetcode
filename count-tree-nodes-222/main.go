package main

import "fmt"

func main() {
	a := &TreeNode{1, &TreeNode{2, &TreeNode{1, nil, nil}, nil}, nil}

	fmt.Println(countNodes(a))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	var (
		ans int
		dfs func(root *TreeNode)
	)

	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		ans++
		dfs(root.Right)
	}

	dfs(root)

	return ans
}

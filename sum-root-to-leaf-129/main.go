package main

import "fmt"

func main() {
	a := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	b := &TreeNode{4, &TreeNode{9, &TreeNode{5, nil, nil}, &TreeNode{1, nil, nil}}, &TreeNode{0, nil, nil}}
	fmt.Printf("%d\n%d\n", sumNumbers(a), sumNumbers(b))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 0
	var dfs func(path []int, root *TreeNode)
	dfs = func(path []int, root *TreeNode) {
		if root.Left == nil && root.Right == nil {
			var val int
			for _, v := range path {
				val = val*10 + v
			}
			ans += val*10 + root.Val
			return
		}

		for _, v := range []*TreeNode{root.Left, root.Right} {
			if v != nil {
				dfs(append(path, root.Val), v)
			}
		}
	}

	dfs(nil, root)

	return ans
}

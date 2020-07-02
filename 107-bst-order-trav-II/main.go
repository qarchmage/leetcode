package main

import "fmt"

func main() {
	r := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}

	fmt.Println(levelOrderBottom(r))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) (res [][]int) {
	res = [][]int{}
	if root == nil {
		return
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		level := []int{}
		for _, x := range q {
			q = q[1:]
			if x.Left != nil {
				q = append(q, x.Left)
			}
			if x.Right != nil {
				q = append(q, x.Right)
			}
			level = append(level, x.Val)
		}
		res = append(res, level)
	}
	for l, r := 0, len(res)-1; l < r; {
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}
	return
}

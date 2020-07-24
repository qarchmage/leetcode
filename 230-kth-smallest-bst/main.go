package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var globalSlice []int

func main() {
	d := &TreeNode{2, nil, nil}
	c := &TreeNode{1, nil, nil}
	b := &TreeNode{5, nil, d}
	a := &TreeNode{2, b, c}
	fmt.Println(kthSmallest(a, 2))
	fmt.Println(globalSlice)

}

// this works but not on leetcode plattform

func kthSmallest(r *TreeNode, k int) int {
	r.addR()

	return globalSlice[k-1]

}

func (t *TreeNode) addR() {
	if t == nil {
		return
	}
	t.Left.addR()
	globalSlice = append(globalSlice, t.Val)
	t.Right.addR()
}

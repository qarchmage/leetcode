package main

import (
	"fmt"
	"sort"
)

// leetcode doesnt like global variables for simplicity i used one
// otherwise u need to pass it around which will allocate memory each time

var ans = make([]int, 0)

func main() {
	a := &TreeNode{0, &TreeNode{-10, nil, nil}, &TreeNode{-10, nil, nil}}
	b := &TreeNode{5, &TreeNode{1, nil, nil}, &TreeNode{7, &TreeNode{2, nil, nil}, &TreeNode{0, nil, nil}}}
	fmt.Println(getAllElements(a, b))
}

// TreeNode  binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//Add all elements from both tree to arr
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	WalkRecursive(root1)
	WalkRecursive(root2)
	sort.Ints(ans)
	return ans
}

// WalkRecursive visit all nodes
func WalkRecursive(t *TreeNode) {
	if t != nil {
		WalkRecursive(t.Left)
		ans = append(ans, t.Val)
		WalkRecursive(t.Right)
	}
}

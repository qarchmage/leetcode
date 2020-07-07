package main

import (
	"fmt"
	"strconv"
)

// this only works if int doesnt overflow which can happen if numbers are big
func main() {
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	walk(addTwoNumbers(l1, l2))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var ls string
	var dfs func(l *ListNode)
	dfs = func(l *ListNode) {
		if l == nil {
			return
		}
		dfs(l.Next)
		ls += strconv.Itoa(l.Val)
	}

	dfs(l1)
	n1, _ := strconv.Atoi(ls)
	ls = ""
	dfs(l2)
	n2, _ := strconv.Atoi(ls)
	newll := n1 + n2
	root := &ListNode{newll % 10, nil}
	newll /= 10
	next := root
	for newll != 0 {
		newLinked := &ListNode{newll % 10, nil}
		next.Next = newLinked
		newll /= 10
		next = newLinked
	}

	return root
}

func walk(l *ListNode) {
	if l == nil {
		return
	}
	walk(l.Next)
	fmt.Println(l.Val)
}

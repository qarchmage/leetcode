package main

import "fmt"

// ListNode linked list implementation
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	d := &ListNode{9, nil}
	c := &ListNode{1, d}
	b := &ListNode{5, c}
	a := &ListNode{4, b}
	printLinkedList(a)
	deleteNode(b)
	printLinkedList(a)
}

func deleteNode(node *ListNode) {
	// if tail dont do it
	if node.Next == nil {
		return
	}
	*node = *node.Next
}

func printLinkedList(node *ListNode) {
	fmt.Printf("%d ", node.Val)
	if node.Next == nil {
		fmt.Printf("\n")
		return
	}
	printLinkedList(node.Next)
}

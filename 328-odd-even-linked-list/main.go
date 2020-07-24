package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	d := ListNode{4, nil}
	c := ListNode{3, &d}
	b := ListNode{2, &c}
	a := ListNode{1, &b}
	display(a)

	oddEvenList(&a)

	display(a)
}
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	odd := head
	evenHead := odd.Next
	even := evenHead

	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head

}

func display(a ListNode) {
	for {
		// current value
		fmt.Println(a.Val)
		// set to next
		if a.Next != nil {
			a = *a.Next
		} else {
			break
		}
	}

}

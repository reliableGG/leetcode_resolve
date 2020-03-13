package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 3}
	l3 := &ListNode{Val: 4}
	l1.Next = l2
	l2.Next = l3

	l4 := &ListNode{Val: 1}
	l5 := &ListNode{Val: 2}
	l6 := &ListNode{Val: 4}

	l4.Next = l5
	l5.Next = l6
	l := mergeTwoLists(l1, l4)
	for l != nil {
		fmt.Printf("%d -> ", l.Val)
		l = l.Next
	}
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	p := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}
	if l1 == nil {
		p.Next = l2
	}
	if l2 == nil {
		p.Next = l1
	}
	return dummy.Next
}

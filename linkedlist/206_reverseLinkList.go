package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	fmt.Println("vim-go")
	l1 := new(ListNode)
	l2 := new(ListNode)
	l3 := new(ListNode)
	l4 := new(ListNode)
	l5 := new(ListNode)
	l6 := new(ListNode)
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	l5.Next = l6
	l6.Next = nil
	l1.Val = 1
	l2.Val = 2
	l3.Val = 3
	l4.Val = 4
	l5.Val = 5
	l6.Val = 6

	l := new(ListNode)
	node := reverseList(l)
	// node := l1
	fmt.Println(node.Next)
	for node != nil {
		fmt.Println("node Val is ", node.Val)
		node = node.Next
	}
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var tail, p, q *ListNode
	tail = nil
	p = head
	q = head.Next
	for q != nil {
		p.Next = tail
		tail = p
		p = q
		q = q.Next
	}
	p.Next = tail
	return p
}

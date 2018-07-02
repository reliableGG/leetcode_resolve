package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	p1 := head
	p2 := head.Next
	for p2.Next != nil && p2 != nil {
		fmt.Println("dubug p2.Val => %d", p2.Val)
		p1 = p1.Next
		p2 = p2.Next.Next
	}

	head2 := p1.Next
	p1.Next = nil

	p2 = head2.Next
	head2.Next = nil

	for p2 != nil {
		p1 = p2.Next
		p2.Next = head2
		head2 = p2
		p2 = p1
	}

	p1 = head
	p2 = head2
	n1 := head
	n2 := head2
	for p1 != nil && p2 != nil {
		n1 = p1.Next
		n2 = p2.Next

		p1.Next = p2
		p2.Next = n1

		p1 = n1
		p2 = n2
	}
}

func main() {
	fmt.Println("vim-go")
	l1 := new(ListNode)
	l2 := new(ListNode)
	l3 := new(ListNode)
	l4 := new(ListNode)
	l5 := new(ListNode)
	l6 := new(ListNode)
	l1.Val = 1
	l2.Val = 2
	l3.Val = 3
	l4.Val = 4
	l5.Val = 5
	l6.Val = 6
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	l5.Next = l6
	l6.Next = nil
	reorderList(l1)
	node := l1
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	p1 := dummy
	p2 := p1
	for p1 != nil && p1.Next != nil {
		p2 = p1.Next
		if p1.Next.Val == val {
			for p2 != nil && p2.Val == val {
				p2 = p2.Next
			}
		}
		p1.Next = p2
		p1 = p1.Next
	}
	return dummy.Next
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
	l4.Val = 3
	l5.Val = 6
	l6.Val = 6
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	l5.Next = l6
	l6.Next = nil
	node := removeElements(l1, 3)
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

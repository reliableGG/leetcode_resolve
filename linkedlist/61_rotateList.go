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
	//l4 := new(ListNode)
	//l5 := new(ListNode)
	//l6 := new(ListNode)
	l1.Val = 0
	l2.Val = 1
	l3.Val = 2
	//l4.Val = 4
	//l5.Val = 8
	//l6.Val = 6
	l1.Next = l2
	l2.Next = l3
	l3.Next = nil
	//l4.Next = nil
	//l4.Next = l5
	//l5.Next = l6
	//l6.Next = nil
	node := rotateRight(l1, 4)
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	h := head
	length := 1
	for h.Next != nil {
		h = h.Next
		length++
	}
	k %= length
	h.Next = head
	p := h
	for i := 0; i < length-k; i++ {
		p = p.Next
		head = head.Next
	}
	p.Next = nil
	return head
}

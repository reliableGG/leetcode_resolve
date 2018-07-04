package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummyhead := new(ListNode)
	dummyhead.Next = head
	p := dummyhead
	p1 := dummyhead
	//for head.Next != nil && head.Next.Next != nil {
	for p1 != nil && p1.Next != nil {
		p1 = head.Next.Next
		head.Next.Next = head
		p.Next = head.Next
		head.Next = p1
		p = head
		head = head.Next
	}
	return dummyhead.Next
}

func main() {
	fmt.Println("vim-go")
	l1 := new(ListNode)
	l2 := new(ListNode)
	l3 := new(ListNode)
	//l4 := new(ListNode)
	//l5 := new(ListNode)
	//l6 := new(ListNode)
	l1.Val = 1
	l2.Val = 2
	l3.Val = 3
	//l4.Val = 4
	//l5.Val = 5
	//l6.Val = 6
	l1.Next = l2
	l2.Next = l3
	l3.Next = nil
	//l4.Next = nil
	//l5.Next = l6
	//l6.Next = nil
	node := swapPairs(l1)
	//node := l1
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	dummyhead := head
	prev := head
	for head != nil {

		for head.Next != nil && head.Val == head.Next.Val {
			head = head.Next
		}
		prev.Next = head.Next
		prev = prev.Next
		head = head.Next
	}
	return dummyhead
}

func main() {
	l1 := new(ListNode)
	l2 := new(ListNode)
	l3 := new(ListNode)
	l4 := new(ListNode)
	l5 := new(ListNode)
	l6 := new(ListNode)
	l7 := new(ListNode)
	l1.Val = 1
	l2.Val = 2
	l3.Val = 3
	l4.Val = 3
	l5.Val = 4
	l6.Val = 4
	l7.Val = 5
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	l5.Next = l6
	l6.Next = l7
	l7.Next = nil

	node := deleteDuplicates(l1)

	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	dummyhead := new(ListNode)
	dummyhead.Next = head
	tail := dummyhead
	pre := dummyhead
	for i := 0; i < m-1; i++ {
		pre = pre.Next
	}
	tail = pre.Next
	temp := tail
	for j := 0; j < n-m; j++ {
		temp = pre.Next
		pre.Next = tail.Next
		tail.Next = tail.Next.Next
		pre.Next.Next = temp
	}
	return dummyhead.Next
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

	node := reverseBetween(l1, 2, 5)
	// node := l1
	for node != nil {
		fmt.Println("node Val is %d", node.Val)
		node = node.Next
	}
}

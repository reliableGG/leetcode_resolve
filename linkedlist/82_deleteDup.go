package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeDupII(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := new(ListNode)
	dummy.Next = head
	p := dummy
	q := p.Next
	fmt.Println(q)
	for q != nil {
		dup := false
		for q.Next != nil && q.Val == q.Next.Val {
			q = q.Next
			dup = true
		}

		if dup {
			q = q.Next
			continue
		}
		p.Next = q
		p = p.Next
		q = q.Next
	}
	p.Next = q
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
	l2.Val = 1
	l3.Val = 2
	l4.Val = 3
	l5.Val = 6
	l6.Val = 6
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	l5.Next = l6
	l6.Next = nil
	node := removeDupII(l1)
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

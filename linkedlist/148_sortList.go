package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head

	if head == nil || head.Next == nil {
		return head
	}
	p := head
	n := head
	nn := head

	for head != nil && n != nil && nn != nil {
		n = head.Next
		nn = n.Next
		p = dummy
		for p != head.Next {

			if n.Val < p.Next.Val {
				n.Next = p.Next
				p.Next = n
				head.Next = nn
				break
			}
			p = p.Next
			if p == head {
				head = head.Next
				break
			}
		}
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
	l2.Val = 7
	l3.Val = 4
	l4.Val = 5
	l5.Val = 8
	l6.Val = 6
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	l5.Next = l6
	l6.Next = nil
	node := sortList(l1)
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

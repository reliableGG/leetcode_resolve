package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	if head.Next == nil {
		return false
	}

	p1 := head
	p2 := head.Next

	for p2 != nil && p2.Next != nil {
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

	for head != nil && head2 != nil {
		if head.Val == head2.Val {
			head = head.Next
			head2 = head2.Next
		} else {
			return false
		}
	}
	return true
}

func main() {
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
	l5.Val = 2
	l6.Val = 1
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	l5.Next = l6
	l6.Next = nil
	res := isPalindrome(l1)
	fmt.Println(res)
}

package main

import "fmt"

func main() {
	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 3}
	l3 := &ListNode{Val: 4}
	l4 := &ListNode{Val: 5}
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = nil
	l := swapPairs(l1)
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}

}

type ListNode struct {
	Val  int
	Next *ListNode
}

//func swapPairs(head *ListNode) *ListNode {
//	dummy := &ListNode{}
//	dummy.Next = head
//
//	if head == nil || head.Next == nil {
//		return head
//	}
//	pre, p1 := dummy, head.Next
//
//	for p1 != nil {
//		head.Next = p1.Next
//		p1.Next = head
//		pre.Next = p1
//
//		if head.Next == nil {
//			break
//		}
//		pre = head
//		head = head.Next
//		p1 = head.Next
//	}
//	return dummy.Next
//}

func swapPairs(head *ListNode) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head

	if head == nil || head.Next == nil {
		return head
	}
	pre := dummy
	tail := dummy.Next

	for tail != nil && tail.Next != nil {
		pre.Next = tail.Next
		tail.Next = tail.Next.Next
		pre.Next.Next = tail

		pre = tail
		tail = tail.Next
	}

	return dummy.Next

}

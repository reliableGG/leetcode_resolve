package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

//func deleteDuplicates(head *ListNode) *ListNode {
//	if head == nil {
//		return nil
//	}
//	//dummyhead := new(ListNode)
//	dummyhead := new(ListNode)
//	dummyhead.Next = head
//	prev := dummyhead
//	mid := dummyhead.Next
//	next := mid.Next
//	curr := dummyhead
//
//	for next != nil {
//		if mid.Val != prev.Val && mid.Val != next.Val {
//			curr.Next = mid
//			curr = curr.Next
//			curr.Next = nil
//		}
//
//		prev = prev.Next
//		mid = prev.Next
//		next = mid.Next
//
//	}
//	if mid.Next == nil && mid.Val != prev.Val {
//		curr.Next = mid
//	}
//	return dummyhead.Next
//}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	for pre, cur := head, head.Next; cur != nil; cur = cur.Next {
		if pre.Val == cur.Val {
			pre.Next = cur.Next
		} else {
			pre = cur
		}
	}
	return head
}

func main() {
	l1 := &ListNode{1, nil}
	l2 := &ListNode{1, nil}
	//l3 := &ListNode{3, nil}
	//l4 := &ListNode{3, nil}
	//l5 := &ListNode{4, nil}
	//l6 := &ListNode{4, nil}
	//l7 := &ListNode{5, nil}

	l1.Next = l2
	l2.Next = nil
	//l3.Next = l4
	//l4.Next = l5
	//l5.Next = l6
	//l6.Next = l7

	node := deleteDuplicates(l1)

	// node := l1
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

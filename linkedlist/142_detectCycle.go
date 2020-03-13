package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	next := head
	for head != nil {
		next = head.Next
		for next != nil {
			fmt.Println("head val is %d", head.Val)
			fmt.Println("next val is %d", next.Val)
			if next.Next != head {
				next = next.Next
			} else {
				return head
			}
		}
		head = head.Next
	}
	return nil
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
	l5.Val = 5
	l6.Val = 6
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	l5.Next = l3
	l6.Next = nil
	node := detectCycle(l1)
	fmt.Println("return is ", node.Val)
}

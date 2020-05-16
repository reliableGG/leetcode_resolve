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
	l4 := new(ListNode)
	l5 := new(ListNode)
	l6 := new(ListNode)
	l1.Val = 1
	l2.Val = 4
	l3.Val = 3
	l4.Val = 2
	l5.Val = 2
	l6.Val = 5
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	l5.Next = l6
	l6.Next = nil
	node := partition(l1, 3)
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

func partition(head *ListNode, x int) *ListNode {
	leftDummy := new(ListNode)
	rightDummy := new(ListNode)

	left := leftDummy
	right := rightDummy

	for cur := head; cur != nil; cur = cur.Next {
		if cur.Val < x {
			left.Next = cur
			left = cur
		} else {
			right.Next = cur
			right = cur
		}
	}

	/*for head != nil {*/
	//tmp := head.Next
	//if head.Val < x {
	//left.Next = head
	//left = left.Next
	//} else {
	//right.Next = head
	//right = right.Next
	//}

	//head.Next = nil
	//head = tmp
	/*}*/

	left.Next = rightDummy.Next
	right.Next = nil

	return leftDummy.Next
}

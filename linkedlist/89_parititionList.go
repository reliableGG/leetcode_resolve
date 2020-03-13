package main

import "fmt"

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
	node := partition(l1)
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

func partition(head *ListNode, x int) *ListNode {

}

package main

import "fmt"

//import "time"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

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
	node := oddEvenList(l1)
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}
	ph1 := head
	ph2 := head.Next
	p1 := head
	p2 := head.Next
	for p1 != nil && p2 != nil && p2.Next != nil {
		p1.Next = p1.Next.Next
		p1 = p1.Next
		p2n := p1.Next
		p2.Next = p2n
		p2 = p2.Next
	}

	//for p1.Next != nil && p1.Next.Next != nil {
	//	p1.Next = p1.Next.Next
	//	p1 = p1.Next
	//}
	//fmt.Println(p1.Val)
	//time.Sleep(time.Second * 2)

	//for p2.Next != nil && p2.Next.Next != nil {
	//	p2.Next = p2.Next.Next
	//	p2 = p2.Next
	//	fmt.Println(p2.Val, "xx")
	//}
	//fmt.Println(p2.Val)
	//time.Sleep(time.Second * 2)
	p1.Next = ph2
	return ph1

}

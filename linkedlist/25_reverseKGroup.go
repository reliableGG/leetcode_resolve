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
	l := reverseKGroup(l1, 3)
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	pre := dummy
	end := head

	for end != nil {
		for i := 1; i < k && end != nil; i++ {
			end = end.Next
			if end == nil {
				break
			}
			pre = reverse(pre, pre.Next, end)
		}
		end = pre.Next
	}
	return dummy.Next
}

func reverse(pre, begin, end *ListNode) *ListNode {
	end_next := end.Next

	p := begin
	cur := p.Next
	next := cur.Next

	for cur != end_next {
		cur.Next = p
		p = cur
		cur = next
		if next != nil {
			next = next.Next
		} else {
			next = nil
		}

	}
	pre.Next = end
	begin.Next = end_next
	return begin
}

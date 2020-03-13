package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dumpHead := head
	length := 0
	search := head
	for search != nil {
		search = search.Next
		length++
	}
	fmt.Println("length is %d", length)
	var preSearch *ListNode = nil

	search = dumpHead
	for i := n; i < length; i++ {
		fmt.Println(i)
		preSearch = search
		search = search.Next
		fmt.Println("search.Val is ", search.Val)
	}
	if preSearch == nil {
		return search.Next
	}
	preSearch.Next = search.Next
	if search.Next == nil {
		preSearch.Next = nil
	}
	return head
}

func main() {
	l1 := new(ListNode)
	l2 := new(ListNode)
	l3 := new(ListNode)
	l1.Next = l2
	l2.Next = l3
	l3.Next = nil
	l1.Val = 1
	l2.Val = 2
	l3.Val = 3
	node := removeNthFromEnd(l1, 3)
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}

}

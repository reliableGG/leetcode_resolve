package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	dummyhead := new(ListNode)
	dummyhead.Next = head
	tail := dummyhead
	pre := dummyhead
	for i := 0; i < m-1; i++ {
		pre = pre.Next
	}
	tail = pre.Next
	temp := tail
	for j := 0; j < n-m; j++ {
		temp = p.Next
		p.Next = tail.Next
		tail.Next = tail.Next.Next
		tail.Next.Next = temp
	}
	return dummyhead.Next
}







}
func main() {
	fmt.Println("vim-go")
}

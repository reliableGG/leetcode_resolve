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

	length := 0

	for i := 0; head != nil; i++ {
		head = head.Next
		length++
	}

	stack := make([]int, length)

	for i := 0; head != nil; i++ {
		stack[i] = head.Val
	}

	mid := length/2 + length%2
	if length%2 > 0 {
		for i := 0; i < mid; i++ {
			if stack[mid-1] == stack[mid+1] {
				continue
			} else {
				return false
			}
		}
		return true
	} else {
		for i := 0; i < mid; i++ {
			if stack[mid-1] == stack[mid] {
				continue
			} else {
				return false
			}
		}
		return true
	}
	return true
}

func main() {
	l1 := new(ListNode)
	l2 := new(ListNode)
	l1.Val = 1
	l1.Next = l2
	l2.Val = 2
	l2.Next = nil
	res := isPalindrome(l1)
	fmt.Println(res)
}

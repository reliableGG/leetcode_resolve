func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil      { return nil }
	if head.Next == nil { return head }

	p := head.Next
	for p != nil && p.Val == head.Val {
		p = p.Next
	}
	if p = head.Next {
		head.Next = deleteDuplicates(p)
		return head
	} else {
		return deleteDuplicates(p)
	}
}


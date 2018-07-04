package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	pq := make(PriorityQueue, len(lists))
	for i := range lists {
		pq[i] = lists[i]
	}
	heap.Init(&pq)
	if pq.Len() < 1 {
		return nil
	}

	result := heap.Pop(&pq).(*ListNode)
	if result == nil {
		return result
	}
	if result.Next != nil {
		heap.Push(&pq, result.Next)
	}
	tail := result
	for pq.Len() > 0 {
		tail.Next = heap.Pop(&pq).(*ListNode)
		tail = tail.Next
		if tail == nil {
			break
		}
		if tail.Next != nil {
			heap.Push(&pq, tail.Next)
		}
	}

	return result
}

type PriorityQueue []*ListNode

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i] == nil {
		return false
	}
	if pq[j] == nil {
		return true
	}
	return pq[i].Val < pq[j].Val
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*ListNode)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func main() {
	fmt.Println("vim-go")
}

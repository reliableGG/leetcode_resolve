package main

import (
	"container/heap"
	"fmt"
)

type KthLargest struct {
	queue     []int
	lockIndex int
}

func (k KthLargest) Less(i, j int) {
	return k.queue[i] > k.queue[j]
}
func (k KthLargest) Len() int {
	return len(k.queue)
}

func (k KthLargest) Swap(i, j int) {
	k.queue[i], k.queue[j] = k.queue[j], k.queue[i]
}

func (k
func Constructor(k int, nums []int) KthLargest {

	return KthLargest{queue: nums, lockIndex: k}
}

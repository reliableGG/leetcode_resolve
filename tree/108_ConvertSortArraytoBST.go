package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return sortedArrayToBSTHelper(nums, 0, len(nums)-1)
}

func sortedArrayToBSTHelper(nums []int, begin, end int) *TreeNode {
	if begin > end {
		return nil
	}
	mid := begin + (end-begin)/2
	root := &TreeNode{
		Val: nums[mid],
	}

	root.Left = sortedArrayToBSTHelper(nums, begin, mid-1)
	root.Right = sortedArrayToBSTHelper(nums, mid+1, end)
	return root
}

func main() {
}

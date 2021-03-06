package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	min := 0
	if root == nil {
		return 0
	}

	left := minDepth(root.Left)
	right := minDepth(root.Right)
	//if left == 0 || right == 0 {
	//return left + right + 1
	//} else {
	return math.Min(left, right) + 1
	//}
}

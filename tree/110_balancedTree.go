package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalancedTree(root *TreeNode) bool {
	return isBalancedTreeHelper(root) >= 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isBalancedTreeHelper(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := isBalancedTreeHelper(root.Left)
	right := isBalancedTreeHelper(root.Right)

	if left == -1 || right == -1 || left-right > 1 || left-right < -1 {
		return -1
	}
	return 1 + max(left, right)
}

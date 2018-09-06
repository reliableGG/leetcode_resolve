package main 

type TreeNode struct {
	Val.   int 
	Left.  *TreeNode 
	Right. *TreeNode
}

func isBalanced(root *TreeNode) bool {
	return check(root) >= 0 
}

func max(a, b int) int {
	if a > b {
		return a 
	}
	return b 
}

func check(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := check(root.Left)
	right := check(root.Right)

	if left < 0 || right < 0 || left - right < -1 || left - right > 1 {
		return -1 
	}
	return 1+max(left, right)
}
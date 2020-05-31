package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
}

func sum(root *TreeNode, su int) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return su*10 + root.Val
	}
	return sum(root.Left, su*10+root.Val) + sum(root.Right, su*10+root.Val)
}

func main() {

}

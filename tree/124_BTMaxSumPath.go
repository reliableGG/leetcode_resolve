package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var maxSum int = 0

func maxPathSum(root *TreeNode) int {
	dfs(root)
	return maxSum
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := dfs(root.Left)
	right := dfs(root.Right)
	sum := root.Val
	if left > 0 {
		sum += left
	}
	if right > 0 {
		sum += right
	}
	maxSum = max(sum, maxSum)
	if max(left, right) > 0 {
		return root.Val + max(left, right)
	}
	return root.Val
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

	n1 := TreeNode{Val: 1}
	n2 := TreeNode{Val: 2}
	n3 := TreeNode{Val: 3}
	n4 := TreeNode{Val: 4}
	n5 := TreeNode{Val: 5}
	n6 := TreeNode{Val: 6}

	n1.Left = &n2
	n1.Right = &n5
	n2.Left = &n3
	n2.Right = &n4
	n3.Left = nil
	n3.Right = nil
	n4.Left = nil
	n4.Right = nil
	n5.Left = nil
	n5.Right = &n6
	n6.Left = nil
	n6.Right = nil

	n7 := new(TreeNode)

	result := maxPathSum(n7)
	fmt.Println(result)
}

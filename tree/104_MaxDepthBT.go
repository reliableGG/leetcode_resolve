package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 *func maxDepth(root *TreeNode) int {
 *    if root == nil {
 *        return 0
 *    }
 *    return maxDepth(root.Left) + 1
 *    return maxDepth(root.Right) + 1
 *}
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lDepth := maxDepth(root.Left)
	rDepth := maxDepth(root.Right)
	return max(lDepth, rDepth) + 1
}

func max(a int, b int) int {
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

	result := maxDepth(&n1)
	fmt.Println(result)
}

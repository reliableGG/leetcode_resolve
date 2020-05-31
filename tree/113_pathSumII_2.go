package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	var slice [][]int
	findPath(root, sum, &slice, []int{})
	return slice
}
func findPath(node *TreeNode, sum int, slice *[][]int, path []int) {
	if node == nil {
		return
	}
	path = append(path, node.Val)
	if node.Left == nil && node.Right == nil {
		if sum == node.Val {
			*slice = append(*slice, path)
		}
	}

	for _, child := range []*TreeNode{node.Left, node.Right} {
		findPath(child, sum-node.Val, slice, path)
	}
}

func main() {
	n1 := TreeNode{Val: 3}
	n2 := TreeNode{Val: 9}
	n3 := TreeNode{Val: 20}
	n4 := TreeNode{Val: 7}
	n5 := TreeNode{Val: 15}

	n1.Left = &n2
	n1.Right = &n3
	n2.Left = nil
	n2.Right = nil
	n3.Left = &n4
	n3.Right = &n5
	n4.Left = nil
	n4.Right = nil
	n5.Left = nil
	n5.Right = nil

	result := pathSum(&n1, 30)
	fmt.Printf("result => %v\n", result)
}

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	var slice [][]int
	findPath(root, sum, &slice, []int(nil))
	return slice
}

func findPath(node *TreeNode, sum int, slice *[][]int, stack []int) {
	if node == nil {
		return
	}
	sum -= node.Val
	stack = append(stack, node.Val)

	if sum == 0 && node.Left == nil && node.Right == nil {
		*slice = append(*slice, append([]int(nil), stack...))
		stack = stack[:len(stack)-1]
	}

	findPath(node.Left, sum, slice, stack)
	findPath(node.Right, sum, slice, stack)
}

func main() {
	n1 := TreeNode{Val: 3}
	n2 := TreeNode{Val: 9}
	n3 := TreeNode{Val: 20}
	n4 := TreeNode{Val: 15}
	n5 := TreeNode{Val: 7}

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

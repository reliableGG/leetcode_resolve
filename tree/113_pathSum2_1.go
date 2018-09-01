package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	var slice [][]int
	slice = findPath(root, sum, slice, []int(nil))
	return slice
}

func findPath(n *TreeNode, sum int, slice [][]int, stack []int) [][]int {
	if n == nil {
		return slice
	}
	sum -= n.Val
	stack = append(stack, n.Val)

	if sum == 0 && n.Left == nil && n.Right == nil {
		slice = append(slice, append([]int(nil), stack...))
		stack = stack[:len(stack)-1]
		fmt.Printf("stack => %v\n", stack)
	}

	slice = findPath(n.Left, sum, slice, stack)
	slice = findPath(n.Right, sum, slice, stack)
	return slice
}

func main() {
	n1 := TreeNode{Val: 5}
	n2 := TreeNode{Val: 4}
	n3 := TreeNode{Val: 8}
	n4 := TreeNode{Val: 11}
	n5 := TreeNode{Val: 13}
	n6 := TreeNode{Val: 4}

	n1.Left = &n2
	n1.Right = &n3
	n2.Left = &n4
	n2.Right = nil
	n3.Left = &n5
	n3.Right = &n6
	n4.Left = nil
	n4.Right = nil
	n5.Left = nil
	n5.Right = nil
	n6.Left = nil
	n6.Right = nil

	res := pathSum(&n1, 20)
	fmt.Printf("%v\n", res)
}

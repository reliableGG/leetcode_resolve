package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	stack := []*TreeNode{}
	var result []int

	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		if len(stack) == 0 {
			break
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append(result, node.Val)

		if node.Right != nil {
			root = node.Right
		}
	}
	return result
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

	node := inorderTraversal(&n1)
	fmt.Println(node)
}

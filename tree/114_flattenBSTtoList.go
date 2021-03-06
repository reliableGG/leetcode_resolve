package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)

	if root.Left == nil {
		return
	}

	p := root.Left
	for p.Right != nil {
		p = p.Right
	}
	p.Right = root.Right
	root.Right = root.Left
	root.Left = nil
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

	flatten(&n1)

	dummy := &n1

	for dummy != nil {
		fmt.Printf("%v -> ", dummy.Val)
		dummy = dummy.Right
	}
}

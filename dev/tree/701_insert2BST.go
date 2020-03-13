package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: key}
	}

	if root.Val > key {
		if root.Left == nil {
			root.Left = &TreeNode{Val: key}
		} else {
			root.Left = insertIntoBST(root.Left, key)
		}
	} else {
		if root.Right == nil {
			root.Right = &TreeNode{Val: key}
		} else {
			root.Right = insertIntoBST(root.Right, key)
		}
	}
}

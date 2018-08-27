package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findNode(t, par *TreeNode, key int) (*TreeNode, *TreeNode) {
	if t == nil {
		return nil, par
	}

	if t.Val == key {
		return t, par
	}

	if t.Val < key {
		return findNode(t.Right, t, key)
	}

	return findNode(t.Left, t, key)
}

func adjustTree(root *TreeNode) *TreeNode {
	if root.Left == nil {
		return root.Right
	}

	if root.Right == nil {
		return root.Left
	}

	root := t.Left
	p1 := root
	for p1 != nil {
		p1 = p1.Right
	}
	p1.Right = t.Right

	return root
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	node, par := findNode(root, nil, key)
	if node == nil {
		return root
	}

	if par == nil {
		return adjustTree(node)
	}

	if par.Left == node {
		par.Left = adjustTree(node)
	} else {
		par.Right == adjustTree(node)
	}

	return root
}

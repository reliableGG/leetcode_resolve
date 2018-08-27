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

func adjustTree(t *TreeNode) *TreeNode {
	if t.Left == nil {
		return t.Right
	}

	if t.Right == nil {
		return t.Left
	}

	root := t.Left
	p1 := root.Right
	for p1 != nil {
		p1 = p1.Right
	}

	p1.Right = t.Right

	return t
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
		par.Right = adjustTree(node)
	}

	return root
}

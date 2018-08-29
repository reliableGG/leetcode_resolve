package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}

	if root.Val == key {
		return merge(root)
	}

	if root.Val < key {
		root.Left = deleteNode(root.Left, key)
	} else {
		root.Right = deleteNode(root.Right, key)
	}

	return root
}

func merge(root *TreeNode) *TreeNode {
	right := root.Right
	left := root.Left

	if right == nil {
		return left
	}

	for right.Left != nil {
		right = right.Left
	}

	right.Left = left
	return root.Right
}

func traval(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("=> %v\n", root.Val)
	traval(root.Left)
	traval(root.Right)
}

func main() {

	n1 := TreeNode{Val: 5}
	n2 := TreeNode{Val: 3}
	n3 := TreeNode{Val: 2}
	n4 := TreeNode{Val: 4}
	n5 := TreeNode{Val: 6}
	n6 := TreeNode{Val: 7}

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

	node := deleteNode(&n1, 5)
	traval(node)
}

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	fmt.Printf("root address => %v\n", root)
	if root == nil || (root.Left == nil && root.Right == nil) {
		return
	}
	flatten(root.Left)
	fmt.Println("root.Left end")
	flatten(root.Right)
	fmt.Printf("root       => %v\n", root)
	fmt.Printf("root.Right => %v\n", root.Right)
	fmt.Printf("root.Left  => %v\n", root.Left)
	currRight := root.Right
	root.Right = root.Left
	root.Left = nil
	/*
	 *for root.Right != nil {
	 *    root = root.Right
	 *}
	 */
	root.Right = currRight
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

	flatten(&n1)
	node := &n1
	for node != nil {
		fmt.Printf("%v\n", node)
		node = node.Right
	}
}

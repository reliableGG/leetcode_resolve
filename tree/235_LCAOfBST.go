package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if max(p.Val, q.Val) < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else if min(p.Val, q.Val) > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	} else {
		return root
	}
}

func lowestCommonAncestor_Iter(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	for root != nil {
		if max(p.Val, q.Val) < root.Val {
			root = root.Left
		} else if min(p.Val, q.Val) > root.Val {
			root = root.Right
		} else {
			return root
		}
	}
	return root
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {

	n1 := TreeNode{Val: 4}
	n2 := TreeNode{Val: 3}
	n3 := TreeNode{Val: 5}
	n4 := TreeNode{Val: 1}
	n5 := TreeNode{Val: 2}
	n6 := TreeNode{Val: 6}

	n1.Left = &n2
	n1.Right = &n3
	n2.Left = &n4
	n2.Right = &n5
	n3.Left = nil
	n3.Right = &n6
	n4.Left = nil
	n4.Right = nil
	n5.Left = nil
	n5.Right = nil
	n6.Left = nil
	n6.Right = nil

	node := lowestCommonAncestor(&n1, &n2, &n3)
	fmt.Println(node)
}

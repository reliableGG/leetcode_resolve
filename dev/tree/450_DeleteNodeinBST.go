package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	t, par := findNode(root, key)
	//fmt.Println(t, par)
	if t == nil {
		return root
	}
	if t == root {
		return merge(t)
	}

	if par.Left == t {
		par.Left = merge(t)
	} else {
		par.Right = merge(t)
	}
	return root
}

func findNode(root *TreeNode, key int) (t, par *TreeNode) {
	par = nil
	t = root

	fmt.Println(key)
	for t != nil {
		fmt.Println("t !=nil ", t)
		if t.Val == key {
			return t, par
		}
		if t != nil && t.Val > key {
			fmt.Println(">")
			fmt.Println(t.Val)
			par = t
			t = t.Left
		}
		if t != nil && t.Val < key {
			par = t
			t = t.Right
		}
	}
	return nil, nil
}

func merge(t *TreeNode) *TreeNode {
	if t == nil || t.Left == nil && t.Right == nil {
		return nil
	}
	if t.Left == nil {
		return t.Right
	}
	if t.Right == nil {
		return t.Left
	}
	tr := t.Right
	for tr.Left != nil {
		tr = tr.Left
	}
	tr.Left = t.Left
	return t.Right
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

	/*
	 *    n1 := TreeNode{Val: 3}
	 *    n2 := TreeNode{Val: 1}
	 *    n3 := TreeNode{Val: 2}
	 *    n4 := TreeNode{Val: 4}
	 *
	 *    n1.Left = &n2
	 *    n1.Right = &n4
	 *    n2.Left = nil
	 *    n2.Right = &n3
	 *    n3.Left = nil
	 *    n3.Right = nil
	 *    n4.Left = nil
	 *    n4.Right = nil
	 */

	//n10 := TreeNode{Val: 0}

	//fmt.Println(n1)
	node := deleteNode(&n1, 0)
	fmt.Println(node)

	fmt.Println(node.Left)
	fmt.Println(node.Left.Left)
	fmt.Println(node.Right)

}

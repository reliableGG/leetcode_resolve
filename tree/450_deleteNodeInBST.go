package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findNode(t, par *TreeNode, key int) (*TreeNode, *TreeNode) {
	fmt.Printf("t => %v\n", t)
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

// adjustTree deletes the given node and adjusts the subtrees under the given node.
func adjustTree(t *TreeNode) *TreeNode {
	if t.Left == nil {
		return t.Right
	}

	if t.Right == nil {
		return t.Left
	}

	root := t.Left
	cur := root

	for cur.Right != nil {
		cur = cur.Right
	}

	cur.Right = t.Right
	/*
	 *fmt.Printf("t => %v\n", t)
	 *fmt.Printf("t.Left => %v\n", t.Left)
	 *fmt.Printf("root => %v\n", root)
	 *fmt.Printf("t.Right => %v\n", t.Right)
	 */

	return root
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	fmt.Println("root")
	cur, par := findNode(root, nil, key)
	fmt.Printf("cur, par => %v \n", cur)
	if cur == nil {
		return root
	}

	if par == nil {
		return adjustTree(cur)
	}

	if par.Left == cur {
		par.Left = adjustTree(cur)
	} else {
		par.Right = adjustTree(cur)
	}

	return root
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

	node := deleteNode(&n1, 2)
	traval(node)
}

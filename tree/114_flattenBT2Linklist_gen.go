package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	p := root
	for node := range nodeGen(root) {
		if node == p {
			continue
		}
		p.Left = nil
		p.Right = node
		p = node
	}
}

func nodeGen(root *TreeNode) <-chan *TreeNode {
	c := make(chan *TreeNode)

	go func() {
		var walk func(r *TreeNode)
		walk = func(r *TreeNode) {
			if r == nil {
				return
			}
			left, right := r.Left, r.Right
			c <- r
			walk(left)
			walk(right)
		}
		walk(root)
		close(c)
	}()

	return c
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

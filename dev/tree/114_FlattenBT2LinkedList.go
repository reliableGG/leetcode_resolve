package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	p := root
	for p != nil {
		if p.Left == nil {
			p = p.Right
		} else {
			pl := p.Left
			for pl.Right != nil {
				pl = pl.Right
			}
			pl.Right = p.Right
			p.Right = p.Left
			p.Left = nil
			p = p.Right
		}
	}
}
func main() {

	n1 := TreeNode{Val: 5}
	n2 := TreeNode{Val: 4}
	n3 := TreeNode{Val: 8}
	n4 := TreeNode{Val: 11}
	n5 := TreeNode{Val: 13}
	n6 := TreeNode{Val: 4}
	n7 := TreeNode{Val: 7}
	n8 := TreeNode{Val: 2}
	n9 := TreeNode{Val: 5}
	n10 := TreeNode{Val: 1}

	n1.Left = &n2
	n1.Right = &n3
	n2.Left = &n4
	n2.Right = nil
	n3.Left = &n5
	n3.Right = &n6
	n4.Left = &n7
	n4.Right = &n8
	n5.Left = nil
	n5.Right = nil
	n6.Left = &n9
	n6.Right = &n10
	n7.Left = nil
	n7.Right = nil
	n8.Left = nil
	n8.Left = nil
	n9.Left = nil
	n9.Left = nil
	n10.Right = nil
	n10.Right = nil

	node := pathSum(&n1, 22)
	fmt.Println(node)
}

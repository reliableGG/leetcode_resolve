package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return -1
	}
	stack := []*TreeNode{root}
	p := root
	for p != nil || len(stack) > 0 {
		for p != nil {
			stack = append(stack, p)
			p = p.Left
		}
		if len(stack) > 0 {
			poped := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			k--
			if k == 0 {
				return poped.Val
			}

			if poped.Right != nil {
				p = poped.Right
			}
		}
	}
	return -1
}

func main() {

	n1 := TreeNode{Val: 4}
	n2 := TreeNode{Val: 2}
	n3 := TreeNode{Val: 5}
	n4 := TreeNode{Val: 1}
	n5 := TreeNode{Val: 3}
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

	k := 2
	node := kthSmallest(&n1, k)
	fmt.Println(node)
}

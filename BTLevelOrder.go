package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		curLevel := []int{}
		qlen := len(queue)
		for i := 0; i < qlen; i++ {
			node := queue[0]
			queue = queue[1:]
			curLevel = append(curLevel, node.Val)
			for _, child := range []*TreeNode{node.Left, node.Right} {
				if child != nil {
					queue = append(queue, child)
				}
			}
		}
		result = append(result, curLevel)
	}
	return result
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

	node := levelOrder(&n1)
	fmt.Println(node)
}

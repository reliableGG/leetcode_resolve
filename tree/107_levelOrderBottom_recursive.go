package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bfs(s *[][]int, level int, root *TreeNode) {
	if root == nil {
		return
	}

	if len(*s) == level {
		*s = append(*s, []int{})
	}

	(*s)[level] = append((*s)[level], root.Val)
	for _, v := range []*TreeNode{root.Left, root.Right} {
		bfs(s, level+1, v)
	}
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var s [][]int
	bfs(&s, 0, root)
	for i := 0; i < len(s); i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
	return s
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

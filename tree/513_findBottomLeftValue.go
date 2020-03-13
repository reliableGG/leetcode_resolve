package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	var res []int
	findBottomLeftValueHelper(root, 0, &res)
	return res[len(res)-1]
}

func findBottomLeftValueHelper(root *TreeNode, level int, res *[]int) {
	if root == nil {
		return
	}

	if len(*res) == level {
		*res = append(*res, root.Val)
	}
	findBottomLeftValueHelper(root.Left, level+1, res)
	findBottomLeftValueHelper(root.Right, level+1, res)
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

	node := findBottomLeftValue(&n1)
	fmt.Println(node)
}

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return isValidBSTHelper(root, nil, nil)
}

func isValidBSTHelper(node *TreeNode, min, max *int) bool {
	if node == nil {
		return true
	}

	if (min != nil && node.Val <= *min) || (max != nil && node.Val >= *max) {
		fmt.Println("node.Val =>", node.Val)
		fmt.Printf("min      =>%v\n", *min)
		fmt.Printf("max      =>%v\n", max)
		return false
	}

	return isValidBSTHelper(node.Left, min, &node.Val) && isValidBSTHelper(node.Right, &node.Val, max)
}

func main() {

	n1 := TreeNode{Val: 6}
	n2 := TreeNode{Val: 2}
	n3 := TreeNode{Val: 1}
	n4 := TreeNode{Val: 3}
	n5 := TreeNode{Val: 7}
	n6 := TreeNode{Val: 9}

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

	node := isValidBST(&n1)
	fmt.Println(node)
}

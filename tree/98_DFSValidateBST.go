package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	res, v := true, root.Val
	doValidateBST(&res, root.Left, &v, nil)
	doValidateBST(&res, root.Right, nil, &v)
	return res
}

func doValidateBST(res *bool, node *TreeNode, upperBound, lowerBound *int) {
	if node == nil || !*res {
		return
	}
	if (upperBound != nil && node.Val >= *upperBound) ||
		(lowerBound != nil && node.Val <= *lowerBound) {
		fmt.Println("node.VaL =>", node.Val)
		fmt.Println("uppper   =>", *upperBound)
		fmt.Println("lower    =>", *lowerBound)
		*res = false
		return
	}

	v := node.Val
	doValidateBST(res, node.Left, &v, lowerBound)
	doValidateBST(res, node.Right, upperBound, &v)
}

func main() {

	n1 := TreeNode{Val: 6}
	n2 := TreeNode{Val: 2}
	n3 := TreeNode{Val: 1}
	n4 := TreeNode{Val: 9}
	n5 := TreeNode{Val: 7}
	n6 := TreeNode{Val: 3}

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

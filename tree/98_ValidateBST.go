package main

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
	doValidateBST(root.Left, *res, nil, &v)
	doValidateBST(root.Right, *res, &v, nil)
	return res
}

func doValidateBST(root *TreeNode, res *bool, lowerBound, upperBound *int) {
	if node == nil || !*res {
		return
	}

	if (upperBound != nil && node.Val > *upperBound) ||
		(lowerBound != nil && node.Val < *lowerBound) {
		*res = false
		return
	}

	v := node.Val
	doValidateBST(res, node.Left, upperBound, &v)
	doValidateBST(res, node.Right, &v, lowerBound)
}

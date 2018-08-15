package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	result := false
	path := make([]int, 0)
	dfs(root, path, sum, &result)
	return result
}

func dfs(root *TreeNode, path []int, sum int, res *bool) {
	if *res {
		return
	}

	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		var result int
		for _, v := range path {
			result += v
		}
		result += root.Val
		if result == sum {
			*res = true
		}
	}

	for _, v := range []*TreeNode{root.Left, root.Right} {
		if v != nil {
			dfs(v, append(path, root.Val), sum, res)
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

	n1.Left = &n2
	n1.Right = &n3
	n2.Left = &n4
	n2.Right = nil
	n3.Left = &n5
	n3.Right = &n6
	n4.Left = nil
	n4.Right = nil
	n5.Left = nil
	n5.Right = nil
	n6.Left = nil
	n6.Right = nil

	res := hasPathSum(nil, 18)
	fmt.Printf("%v\n", res)
}

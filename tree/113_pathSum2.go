package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	var result [][]int
	path := make([]int, 0)
	if root == nil {
		return nil
	}
	dfs(root, path, sum, &result)
	return result
}

func dfs(root *TreeNode, path []int, sum int, result *[][]int) {
	if root.Left == nil && root.Right == nil {
		path = append(path, root.Val)
		var res int
		for _, v := range path {
			res = res + v
		}
		if res == sum {
			*result = append(*result, path)
		}
	}
	for _, v := range []*TreeNode{root.Left, root.Right} {
		if v != nil {
			dfs(v, append(path, root.Val), sum, result)
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

	res := pathSum(&n1, 20)
	fmt.Printf("%v\n", res)
}

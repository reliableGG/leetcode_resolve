package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, path []int, sum *int) {
	var val int
	if root.Left == nil && root.Right == nil {
		for _, v := range path {
			val = val*10 + v
			return
		}
	}

	for _, v := range []*TreeNode{root.Left, root.Right} {
		if v != nil {
			dfs(v, append(path, root.Val), sum)
		}
	}
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var sum int
	dfs(root, nil, &sum)
	return sum
}

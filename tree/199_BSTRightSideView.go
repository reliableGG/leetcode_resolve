package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	result := make([]int, 0)
	rightView(&result, 0, root)
	return result

}

func rightView(res *[]int, hight int, node *TreeNode) {
	if node == nil {
		return
	}

	if hight == len(res) {
		res = append(&res, node.Val)
	}
	rightView(res, hight+1, node.Right)
	rightView(res, hight+1, node.Left)
}

func rightSideViewIter(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	queue := make([]*TreeNode, 0)
	numLevel := 1
	queue = append(queue, root)
	result := make([]int, 0)

	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		numLevel--
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		if numLevel == 0 {
			result = append(result, node)
			numLevel = len(queue)
		}
	}
	return result
}

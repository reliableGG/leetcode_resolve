package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// recursive
func dfs(res *[][]int, root *TreeNode, level int) {
	if root != nil {
		if len(*res) < level+1 {
			*res = append(*res, []int{})

}

func levelOrderBottom1(root *TreeNode) [][]int {
	var result [][]int

}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	queue := make([]*TreeNode, 0)
	result := make([][]int, 0)
	level := 1
	queue = append(queue, root)
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		level--
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		if level == 0 {
			dataQueue := make([]int, 0)
			fmt.Println(len(queue))
			for _, data := range queue {
				dataQueue = append(dataQueue, data.Val)
			}
			fmt.Printf("dataQueue => %v\n", dataQueue)
			result = append(result, dataQueue)
			level = len(queue)
		}
	}
	return result
}

func main() {
	n1 := TreeNode{Val: 3}
	n2 := TreeNode{Val: 9}
	n3 := TreeNode{Val: 20}
	n4 := TreeNode{Val: 15}
	n5 := TreeNode{Val: 7}

	n1.Left = &n2
	n1.Right = &n3
	n2.Left = nil
	n2.Right = nil
	n3.Left = &n4
	n3.Right = &n5
	n4.Left = nil
	n4.Right = nil
	n5.Left = nil
	n5.Right = nil

	result := levelOrderBottom(&n1)
	fmt.Printf("result => %v\n", result)
}

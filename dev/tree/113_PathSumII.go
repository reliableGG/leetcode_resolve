package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	var result [][]int
	dfs(root, []int{}, sum, &result)
	return result
}

func dfs(root *TreeNode, path []int, sum int, s *[][]int) {

	if root == nil {
		return
	}
	path = append(path, root.Val)
	sum = sum - root.Val
	fmt.Println("sum => ", sum)

	if root.Left == nil && root.Right == nil {
		if sum == 0 {
			*s = append(*s, append([]int(nil), path...))
			fmt.Printf("%p\n", s)
			fmt.Printf("%v\n", *s)
			return
		}
	}
	for _, v := range []*TreeNode{root.Left, root.Right} {
		dfs(v, path, sum, s)
	}
}

func main() {

	n1 := TreeNode{Val: 5}
	n2 := TreeNode{Val: 4}
	n3 := TreeNode{Val: 8}
	n4 := TreeNode{Val: 11}
	n5 := TreeNode{Val: 13}
	n6 := TreeNode{Val: 4}
	n7 := TreeNode{Val: 7}
	n8 := TreeNode{Val: 2}
	n9 := TreeNode{Val: 5}
	n10 := TreeNode{Val: 1}

	n1.Left = &n2
	n1.Right = &n3
	n2.Left = &n4
	n2.Right = nil
	n3.Left = &n5
	n3.Right = &n6
	n4.Left = &n7
	n4.Right = &n8
	n5.Left = nil
	n5.Right = nil
	n6.Left = &n9
	n6.Right = &n10
	n7.Left = nil
	n7.Right = nil
	n8.Left = nil
	n8.Left = nil
	n9.Left = nil
	n9.Left = nil
	n10.Right = nil
	n10.Right = nil

	node := pathSum(&n1, 22)
	fmt.Println(node)
}
